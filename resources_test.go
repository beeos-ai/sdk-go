package beeos

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func testClient(t *testing.T, handler http.HandlerFunc) *Client {
	t.Helper()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-key" {
			t.Errorf("missing API authentication")
		}
		w.Header().Set("Content-Type", "application/json")
		handler(w, r)
	}))
	t.Cleanup(server.Close)
	client, err := NewClient(ClientOptions{APIKey: "test-key", BaseURL: server.URL})
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func taskSnapshot(status TaskStatus) TaskResponse {
	return TaskResponse{Success: true, Data: TaskResponseAllOfData{AgentId: "agent", TaskId: "task", Status: status}}
}

func TestResourceHandlesDoNotCreate(t *testing.T) {
	client := testClient(t, func(http.ResponseWriter, *http.Request) { t.Error("unexpected network request") })
	instance := client.Instance("phone")
	if instance.ID() != "phone" || instance.Mobile == nil || instance.Computer == nil {
		t.Fatal("incomplete instance handle")
	}
	if client.Task("agent", "task").ID() != "task" {
		t.Fatal("wrong task handle")
	}
	if _, err := client.Instance("").Mobile.Screenshot(context.Background()); err == nil {
		t.Fatal("empty ID accepted")
	}
}

func TestMobileAndComputerAtomicRoutes(t *testing.T) {
	type call struct{ path, body string }
	calls := []call{}
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		calls = append(calls, call{r.URL.Path, string(body)})
		_ = json.NewEncoder(w).Encode(InlineObject{Success: true, Data: ActionResult{Ok: true}})
	})
	mobile := client.Instance("phone").Mobile
	computer := client.Instance("desktop").Computer
	ctx := context.Background()
	actions := []struct {
		path string
		run  func() (*InlineObject, error)
	}{
		{"/instances/phone/mobile/tap", func() (*InlineObject, error) { return mobile.Tap(ctx, TapRequest{X: 10, Y: 20}) }},
		{"/instances/phone/mobile/double_tap", func() (*InlineObject, error) { return mobile.DoubleTap(ctx, TapRequest{X: 10, Y: 20}) }},
		{"/instances/phone/mobile/long_press", func() (*InlineObject, error) { return mobile.LongPress(ctx, LongPressRequest{X: 10, Y: 20}) }},
		{"/instances/phone/mobile/drag", func() (*InlineObject, error) { return mobile.Drag(ctx, DragRequest{X1: 1, Y1: 2, X2: 3, Y2: 4}) }},
		{"/instances/phone/mobile/swipe", func() (*InlineObject, error) { return mobile.Swipe(ctx, SwipeRequest{X1: 1, Y1: 2, X2: 3, Y2: 4}) }},
		{"/instances/phone/mobile/scroll", func() (*InlineObject, error) { return mobile.Scroll(ctx, ScrollRequest{Direction: "down"}) }},
		{"/instances/phone/mobile/type", func() (*InlineObject, error) { return mobile.Type(ctx, "test") }},
		{"/instances/phone/mobile/key", func() (*InlineObject, error) { return mobile.Key(ctx, "ENTER") }},
		{"/instances/phone/mobile/press_button", func() (*InlineObject, error) { return mobile.PressButton(ctx, "home") }},
		{"/instances/phone/mobile/open_app", func() (*InlineObject, error) { return mobile.OpenApp(ctx, "com.android.settings") }},
		{"/instances/desktop/computer/click", func() (*InlineObject, error) { return computer.Click(ctx, ClickRequest{X: 10, Y: 20}) }},
		{"/instances/desktop/computer/move", func() (*InlineObject, error) { return computer.Move(ctx, MoveRequest{X: 10, Y: 20}) }},
		{"/instances/desktop/computer/scroll", func() (*InlineObject, error) { return computer.Scroll(ctx, ScrollRequest{Direction: "up"}) }},
		{"/instances/desktop/computer/type", func() (*InlineObject, error) { return computer.Type(ctx, "text") }},
		{"/instances/desktop/computer/key", func() (*InlineObject, error) { return computer.Key(ctx, "CTRL", "L") }},
	}
	for _, action := range actions {
		result, err := action.run()
		if err != nil || !result.Data.Ok {
			t.Fatalf("%s: %v", action.path, err)
		}
		if calls[len(calls)-1].path != "/api/v1"+action.path {
			t.Errorf("wrong route: %v", calls[len(calls)-1])
		}
	}
	if !strings.Contains(calls[14].body, `"keys":["CTRL","L"]`) {
		t.Fatalf("hotkeys were not preserved: %s", calls[14].body)
	}
	if !strings.Contains(calls[9].body, `"app":"com.android.settings"`) {
		t.Fatal("wrong app wire field")
	}
}

func TestInstanceLifecycleAndReady(t *testing.T) {
	var polls int
	var operations []string
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		operations = append(operations, r.Method+" "+r.URL.Path)
		data := InstanceDataDTO{Id: "instance", Models: []string{}, ProviderConfig: map[string]interface{}{}, Status: "running", Connectivity: "online"}
		if r.Method == "GET" {
			polls++
			if polls == 1 {
				data.Connectivity = "offline"
			}
		}
		if r.Method == "DELETE" {
			_ = json.NewEncoder(w).Encode(StringResponse{Success: true, Data: "deleted"})
			return
		}
		_ = json.NewEncoder(w).Encode(InstanceResponse{Success: true, Data: data})
	})
	ctx := context.Background()
	instance, err := client.Instances.Create(ctx, DeployInstanceRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if _, err = instance.WaitReady(ctx, WaitOptions{PollInterval: time.Millisecond}); err != nil {
		t.Fatal(err)
	}
	if polls != 2 {
		t.Fatalf("ready accepted offline instance after %d polls", polls)
	}
	for _, operation := range []func(context.Context) (*InstanceResponse, error){instance.Start, instance.Stop, instance.Restart} {
		if _, err = operation(ctx); err != nil {
			t.Fatal(err)
		}
	}
	if _, err = instance.Destroy(ctx); err != nil {
		t.Fatal(err)
	}
	if operations[len(operations)-1] != "DELETE /api/v1/instances/instance" {
		t.Fatalf("unexpected lifecycle routes: %v", operations)
	}
}

func TestTaskWaitStrictTerminals(t *testing.T) {
	for _, status := range []TaskStatus{TASKSTATUS_COMPLETED, TASKSTATUS_FAILED, TASKSTATUS_CANCELED, TASKSTATUS_TIMEOUT, TASKSTATUS_REJECTED} {
		t.Run(string(status), func(t *testing.T) {
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(taskSnapshot(status)) })
			result, err := client.Task("agent", "task").Wait(context.Background())
			if status == TASKSTATUS_COMPLETED {
				if err != nil {
					t.Fatal(err)
				}
			} else {
				var taskError *TaskExecutionError
				if !errors.As(err, &taskError) || taskError.Snapshot != result {
					t.Fatalf("failed terminal counted as success: %v", err)
				}
			}
			if _, err = client.Task("agent", "task").Wait(context.Background(), WaitOptions{AllowUnsuccessful: true}); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestTaskWaitRejectsMismatchedAndTruncatedSnapshot(t *testing.T) {
	for _, kind := range []string{"mismatched", "truncated"} {
		t.Run(kind, func(t *testing.T) {
			snapshot := taskSnapshot(TASKSTATUS_COMPLETED)
			if kind == "mismatched" {
				snapshot.Data.AgentId = "other"
			} else {
				v := true
				snapshot.Data.Truncated = &v
			}
			client := testClient(t, func(w http.ResponseWriter, r *http.Request) { _ = json.NewEncoder(w).Encode(snapshot) })
			if _, err := client.Task("agent", "task").Wait(context.Background()); err == nil {
				t.Fatal("invalid snapshot accepted")
			}
		})
	}
}

func TestWaitCancelsInflightRequestNotRemoteTask(t *testing.T) {
	var writes int32
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			atomic.AddInt32(&writes, 1)
		}
		<-r.Context().Done()
	})
	_, err := client.Task("agent", "task").Wait(context.Background(), WaitOptions{Timeout: 30 * time.Millisecond})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("deadline did not propagate: %v", err)
	}
	if atomic.LoadInt32(&writes) != 0 {
		t.Fatal("local timeout mutated remote task")
	}
}

func TestBusyWriteIsNotRetried(t *testing.T) {
	count := 0
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		count++
		w.WriteHeader(409)
		_, _ = io.WriteString(w, `{"success":false,"error":{"code":"device_busy","message":"busy"}}`)
	})
	if _, err := client.Instance("phone").Mobile.PressButton(context.Background(), "home"); err == nil {
		t.Fatal("busy write succeeded")
	}
	if count != 1 {
		t.Fatalf("write replayed %d times", count)
	}
}

func TestSSECancellationUnblocksAbandonedConsumer(t *testing.T) {
	closed := make(chan struct{})
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		_, _ = io.WriteString(w, "data: {}\n\n")
		w.(http.Flusher).Flush()
		<-r.Context().Done()
		close(closed)
	})
	ctx, cancel := context.WithCancel(context.Background())
	_, errs := client.TaskEvents(ctx, "agent", "task", 0)
	time.Sleep(10 * time.Millisecond)
	cancel()
	select {
	case <-errs:
	case <-time.After(time.Second):
		t.Fatal("SSE sender leaked after cancellation")
	}
	select {
	case <-closed:
	case <-time.After(time.Second):
		t.Fatal("SSE response not closed")
	}
}

func TestSSEJSONErrorIsNotAnEmptySuccess(t *testing.T) {
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) { _, _ = io.WriteString(w, `{"success":false}`) })
	events, errs := client.TaskEvents(context.Background(), "agent", "task", 0)
	for range events {
		t.Error("JSON was treated as SSE")
	}
	if err := <-errs; err == nil {
		t.Fatal("HTTP200 JSON counted as a good stream")
	}
}

func TestSSEFramingAndLimit(t *testing.T) {
	events := make(chan TaskEvent, 5)
	input := "id: 40\r\ndata: {}\r\n\r\nid: 42\ndata: {}\n\nevent: end\ndata: {}\n\ndata: ignored\n\n"
	if err := scanTaskEvents(strings.NewReader(input), events); err != nil {
		t.Fatal(err)
	}
	if len(events) != 3 {
		t.Fatalf("unexpected frame count %d", len(events))
	}
	if (<-events).ID != "40" || (<-events).ID != "42" {
		t.Fatal("noncontiguous offsets lost")
	}
	if err := scanTaskEvents(strings.NewReader("data: incomplete"), events); err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 {
		t.Fatal("incomplete EOF frame accepted")
	}
	if err := scanTaskEvents(strings.NewReader("data: "+strings.Repeat("x", 4*1024*1024)), events); err == nil {
		t.Fatal("oversized frame accepted")
	}
}

func TestFileDownloadIsolationAndGuards(t *testing.T) {
	var downloads int32
	store := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&downloads, 1)
		if r.Header.Get("Authorization") != "" || r.Header.Get("Cookie") != "" {
			t.Error("API credentials reached object storage")
		}
		w.Header().Set("Content-Type", "application/pdf")
		_, _ = io.WriteString(w, "%PDF-test")
	}))
	defer store.Close()
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		uri := store.URL + "/report"
		if r.URL.Path == "/api/v1/files/wrong" {
			_ = json.NewEncoder(w).Encode(GetFile200Response{Success: true, Data: &FileMetaResponse{FileId: "other", DownloadUrl: &uri}})
			return
		}
		_ = json.NewEncoder(w).Encode(GetFile200Response{Success: true, Data: &FileMetaResponse{FileId: "file", DownloadUrl: &uri}})
	})
	ctx := context.Background()
	if _, err := client.Files.Download(ctx, "file"); err == nil {
		t.Fatal("HTTP download allowed implicitly")
	}
	if atomic.LoadInt32(&downloads) != 0 {
		t.Fatal("invalid URL requested")
	}
	bytes, err := client.Files.Download(ctx, "file", DownloadOptions{AllowInsecureHTTP: true, ExpectedContentType: "application/pdf"})
	if err != nil || string(bytes) != "%PDF-test" {
		t.Fatalf("download failed: %v", err)
	}
	if _, err := client.Files.Download(ctx, "file", DownloadOptions{AllowInsecureHTTP: true, MaxBytes: 3}); err == nil {
		t.Fatal("oversized file accepted")
	}
	count := atomic.LoadInt32(&downloads)
	if _, err := client.Files.Download(ctx, "wrong", DownloadOptions{AllowInsecureHTTP: true}); err == nil {
		t.Fatal("mismatched identity accepted")
	}
	if count != atomic.LoadInt32(&downloads) {
		t.Fatal("mismatched file downloaded")
	}
}

func TestFileRedirectIsNotFollowed(t *testing.T) {
	hits := 0
	store := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { hits++; http.Redirect(w, r, "/next", 302) }))
	defer store.Close()
	client := testClient(t, func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, `{"success":true,"data":{"file_id":"file","download_url":%q}}`, store.URL)
	})
	if _, err := client.Files.Download(context.Background(), "file", DownloadOptions{AllowInsecureHTTP: true}); err == nil {
		t.Fatal("redirect accepted")
	}
	if hits != 1 {
		t.Fatal("redirect followed")
	}
}

func TestFileTransportErrorsDoNotExposeSignedURLs(t *testing.T) {
	err := downloadTransportError(&url.Error{Op: "Get", URL: "https://store.example/file?X-Amz-Signature=secret", Err: context.DeadlineExceeded})
	if strings.Contains(err.Error(), "secret") || strings.Contains(err.Error(), "store.example") {
		t.Fatal("signed URL leaked in download error")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatal("timeout cause was lost")
	}
}
