package beeos

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"mime"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type FilesResource struct{ client *Client }
type DownloadOptions struct {
	MaxBytes            int64
	ExpectedContentType string
	AllowInsecureHTTP   bool
	// MaxAttempts defaults to 3. Set 1 to disable safe object GET retries.
	MaxAttempts int
}

func (f *FilesResource) Get(ctx context.Context, id string) (*GetFile200Response, error) {
	return resourceCall(ctx, id, func(ctx context.Context) (*GetFile200Response, *http.Response, error) {
		return f.client.API.FilesAPI.GetFile(f.client.authContext(ctx), id).Execute()
	})
}

// Download resolves owner-authorized file metadata, then uses a credential-free
// request to object storage. It never follows redirects or retries API writes.
func (f *FilesResource) Download(ctx context.Context, id string, options ...DownloadOptions) ([]byte, error) {
	if len(options) > 1 {
		return nil, fmt.Errorf("expected at most one DownloadOptions value")
	}
	opts := DownloadOptions{}
	if len(options) == 1 {
		opts = options[0]
	}
	if opts.MaxAttempts < 0 || opts.MaxAttempts > 10 {
		return nil, fmt.Errorf("MaxAttempts must be between 1 and 10")
	}
	if opts.MaxAttempts == 0 {
		opts.MaxAttempts = 3
	}
	if opts.MaxBytes < 0 || opts.MaxBytes == int64(^uint64(0)>>1) {
		return nil, fmt.Errorf("invalid MaxBytes")
	}
	if opts.MaxBytes == 0 {
		opts.MaxBytes = 32 * 1024 * 1024
	}
	ctx, cancel := requestContext(ctx)
	defer cancel()
	metadata, err := f.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if metadata == nil || !metadata.Success || metadata.Data == nil || metadata.Data.FileId != id || metadata.Data.GetDownloadUrl() == "" {
		return nil, fmt.Errorf("file response is missing an authorized download URL")
	}
	uri, err := url.Parse(metadata.Data.GetDownloadUrl())
	if err != nil || uri.Hostname() == "" || uri.User != nil || (uri.Scheme != "https" && !(opts.AllowInsecureHTTP && uri.Scheme == "http")) {
		return nil, fmt.Errorf("file download requires an HTTPS URL without embedded credentials")
	}
	if metadata.Data.GetSize() > opts.MaxBytes {
		return nil, fmt.Errorf("file exceeds MaxBytes")
	}
	return downloadFile(ctx, uri.String(), opts, newDownloadClient)
}

func newDownloadClient(http1Only bool) *http.Client {
	// A fresh transport avoids inheriting a globally replaced DefaultTransport,
	// API authentication, cookies, or client certificates. TLS stays verified.
	transport := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		DialContext:         (&net.Dialer{Timeout: 30 * time.Second, KeepAlive: 30 * time.Second}).DialContext,
		ForceAttemptHTTP2:   !http1Only,
		TLSHandshakeTimeout: 10 * time.Second,
		IdleConnTimeout:     30 * time.Second,
	}
	if http1Only {
		transport.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}
	}
	return &http.Client{Transport: transport, CheckRedirect: func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }}
}

// Retry the same signed GET, never metadata resolution or the action that
// created the file. Partial bodies are discarded, not appended to a retry.
func downloadFile(ctx context.Context, uri string, opts DownloadOptions, clientFactory func(bool) *http.Client) ([]byte, error) {
	http1Only := false
	for attempt := 0; attempt < opts.MaxAttempts; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		client := clientFactory(http1Only)
		body, retry, transportFailure, retryAfter, err := downloadFileAttempt(ctx, client, uri, opts)
		client.CloseIdleConnections()
		if err == nil {
			return body, nil
		}
		if !retry || attempt+1 == opts.MaxAttempts {
			return nil, err
		}
		// Some proxies reset negotiated HTTP/2 streams. Retry idempotent file
		// reads over HTTP/1.1 without weakening TLS or redirect protections.
		http1Only = http1Only || transportFailure
		delay := time.Duration(100*(1<<uint(attempt))) * time.Millisecond
		if retryAfter > delay {
			delay = retryAfter
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
	return nil, fmt.Errorf("file download has no attempts configured")
}

func downloadFileAttempt(ctx context.Context, client *http.Client, uri string, opts DownloadOptions) ([]byte, bool, bool, time.Duration, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, false, false, 0, fmt.Errorf("invalid download request")
	}
	response, err := client.Do(req)
	if err != nil {
		retry := retryableDownloadError(err)
		return nil, retry, retry, 0, downloadTransportError(err)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		retry := response.StatusCode == 408 || response.StatusCode == 429 || response.StatusCode == 500 || response.StatusCode == 502 || response.StatusCode == 503 || response.StatusCode == 504
		return nil, retry, false, downloadRetryAfter(response.Header.Get("Retry-After")), fmt.Errorf("file download failed: HTTP %d", response.StatusCode)
	}
	mediaType, _, _ := mime.ParseMediaType(response.Header.Get("Content-Type"))
	if opts.ExpectedContentType != "" && !strings.EqualFold(mediaType, opts.ExpectedContentType) {
		return nil, false, false, 0, fmt.Errorf("downloaded file has an unexpected content type")
	}
	if response.ContentLength > opts.MaxBytes {
		return nil, false, false, 0, fmt.Errorf("file exceeds MaxBytes")
	}
	bytes, err := io.ReadAll(io.LimitReader(response.Body, opts.MaxBytes+1))
	if err != nil {
		retry := retryableDownloadError(err)
		return nil, retry, retry, 0, downloadTransportError(err)
	}
	if int64(len(bytes)) > opts.MaxBytes {
		return nil, false, false, 0, fmt.Errorf("file exceeds MaxBytes")
	}
	return bytes, false, false, 0, nil
}

func retryableDownloadError(err error) bool {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		return false
	}
	for _, transient := range []error{io.EOF, io.ErrUnexpectedEOF, syscall.ECONNRESET, syscall.ECONNREFUSED, syscall.EPIPE, syscall.ETIMEDOUT} {
		if errors.Is(err, transient) {
			return true
		}
	}
	var networkError net.Error
	return errors.As(err, &networkError) && networkError.Timeout()
}

func downloadRetryAfter(value string) time.Duration {
	if seconds, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64); err == nil && seconds >= 0 {
		if seconds > int64((1<<63-1)/time.Second) {
			return time.Duration(1<<63 - 1)
		}
		return time.Duration(seconds) * time.Second
	}
	if at, err := http.ParseTime(value); err == nil && at.After(time.Now()) {
		return time.Until(at)
	}
	return 0
}

func downloadTransportError(err error) error {
	var urlError *url.Error
	if errors.As(err, &urlError) {
		// net/http includes the full presigned URL in url.Error.Error(). Do not
		// expose object credentials in a log or a caller's returned error.
		return fmt.Errorf("file download transport failed: %w", urlError.Err)
	}
	return fmt.Errorf("file download transport failed: %w", err)
}
