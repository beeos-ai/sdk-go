package beeos

import (
	"context"
	"crypto/x509"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
)

type downloadRoundTrip func(*http.Request) (*http.Response, error)

func (f downloadRoundTrip) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

type interruptedDownload struct {
	closed bool
	read   bool
}

func (r *interruptedDownload) Read(p []byte) (int, error) {
	if r.read {
		return 0, syscall.ECONNRESET
	}
	r.read = true
	return copy(p, "partial"), nil
}
func (r *interruptedDownload) Close() error { r.closed = true; return nil }

func TestDownloadRetriesOnlyTheSignedGETAndDiscardsPartialBody(t *testing.T) {
	broken := &interruptedDownload{}
	var attempts []bool
	uri := "https://store.example/file?X-Amz-Signature=private"
	factory := func(http1 bool) *http.Client {
		attempts = append(attempts, http1)
		return &http.Client{Transport: downloadRoundTrip(func(req *http.Request) (*http.Response, error) {
			if req.Method != http.MethodGet || req.URL.String() != uri || req.Header.Get("Authorization") != "" || req.Header.Get("Cookie") != "" {
				t.Fatal("signed read changed or inherited authentication")
			}
			var body io.ReadCloser = io.NopCloser(strings.NewReader("complete"))
			if len(attempts) == 1 {
				body = broken
			}
			return &http.Response{StatusCode: 200, Body: body, Header: make(http.Header), ContentLength: -1}, nil
		})}
	}
	bytes, err := downloadFile(context.Background(), uri, DownloadOptions{MaxAttempts: 3, MaxBytes: 1024}, factory)
	if err != nil || string(bytes) != "complete" || !broken.closed {
		t.Fatalf("partial response was not safely replaced: %q, %v", bytes, err)
	}
	if len(attempts) != 2 || attempts[0] || !attempts[1] {
		t.Fatalf("expected verified HTTPS HTTP/1 fallback after reset: %v", attempts)
	}
}

func TestDownloadHTTPRetryPolicyAndMetadataResolvedOnce(t *testing.T) {
	for _, code := range []int{302, 403, 404, 408, 429, 500, 502, 503, 504} {
		t.Run(fmt.Sprint(code), func(t *testing.T) {
			var calls, metadata int32
			store := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if atomic.AddInt32(&calls, 1) == 1 {
					w.WriteHeader(code)
					return
				}
				_, _ = io.WriteString(w, "complete")
			}))
			defer store.Close()
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
				atomic.AddInt32(&metadata, 1)
				_, _ = fmt.Fprintf(w, `{"success":true,"data":{"file_id":"file","download_url":%q}}`, store.URL)
			})
			_, err := client.Files.Download(context.Background(), "file", DownloadOptions{AllowInsecureHTTP: true})
			retry := code >= 408
			if retry && (err != nil || atomic.LoadInt32(&calls) != 2) {
				t.Fatalf("transient failure not retried: %v", err)
			}
			if !retry && (err == nil || atomic.LoadInt32(&calls) != 1) {
				t.Fatalf("permanent failure retried: %v", err)
			}
			if atomic.LoadInt32(&metadata) != 1 {
				t.Fatal("metadata lookup repeated")
			}
		})
	}
}

func TestDownloadRetryBudgetAndExplicitDisable(t *testing.T) {
	for _, attempts := range []int{1, 3} {
		calls := 0
		factory := func(bool) *http.Client {
			return &http.Client{Transport: downloadRoundTrip(func(*http.Request) (*http.Response, error) {
				calls++
				return nil, syscall.ECONNRESET
			})}
		}
		_, err := downloadFile(context.Background(), "https://store.example/file?secret=yes", DownloadOptions{MaxAttempts: attempts, MaxBytes: 1024}, factory)
		if err == nil || calls != attempts || strings.Contains(err.Error(), "secret") || !errors.Is(err, syscall.ECONNRESET) {
			t.Fatalf("incorrect retry budget or unsafe error: %d %v", calls, err)
		}
	}
}

func TestDownloadRetryAfterCannotOutliveContext(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	calls := 0
	factory := func(bool) *http.Client {
		return &http.Client{Transport: downloadRoundTrip(func(*http.Request) (*http.Response, error) {
			calls++
			return &http.Response{StatusCode: 429, Body: io.NopCloser(strings.NewReader("")), Header: http.Header{"Retry-After": []string{"3600"}}}, nil
		})}
	}
	_, err := downloadFile(ctx, "https://store.example/file", DownloadOptions{MaxAttempts: 3, MaxBytes: 1024}, factory)
	if !errors.Is(err, context.DeadlineExceeded) || calls != 1 {
		t.Fatalf("deadline not enforced: %d %v", calls, err)
	}
}

func TestDownloadDoesNotRetryCertificateOrValidationFailures(t *testing.T) {
	for _, err := range []error{x509.UnknownAuthorityError{}, context.Canceled, context.DeadlineExceeded, errors.New("invalid URL")} {
		if retryableDownloadError(err) {
			t.Fatalf("permanent failure considered retryable: %v", err)
		}
	}
	client := testClient(t, func(http.ResponseWriter, *http.Request) { t.Fatal("invalid options reached API") })
	for _, attempts := range []int{-1, 11} {
		if _, err := client.Files.Download(context.Background(), "file", DownloadOptions{MaxAttempts: attempts}); err == nil {
			t.Fatal("invalid attempts accepted")
		}
	}
}

func TestDownloadTransportIsolationAndTLSVerification(t *testing.T) {
	for _, http1 := range []bool{false, true} {
		client := newDownloadClient(http1)
		transport, ok := client.Transport.(*http.Transport)
		if !ok || transport == http.DefaultTransport || client.Jar != nil || (transport.TLSClientConfig != nil && transport.TLSClientConfig.InsecureSkipVerify) {
			t.Fatal("download does not have an isolated, verifying transport")
		}
		if http1 && (transport.ForceAttemptHTTP2 || transport.TLSNextProto == nil) {
			t.Fatal("fallback may negotiate HTTP/2")
		}
		if !http1 && !transport.ForceAttemptHTTP2 {
			t.Fatal("initial HTTP/2 support removed")
		}
		if client.CheckRedirect(&http.Request{}, nil) != http.ErrUseLastResponse {
			t.Fatal("redirect allowed")
		}
	}
}
