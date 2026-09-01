package beeos

import (
	"bytes"
	"testing"
)

func TestScanTaskEvents(t *testing.T) {
	input := "id: 7\nevent: message\ndata: {\"status\":\"working\"}\n\n"
	events := make(chan TaskEvent, 1)
	if err := scanTaskEvents(bytes.NewBufferString(input), events); err != nil {
		t.Fatal(err)
	}
	event := <-events
	if event.ID != "7" || event.Event != "message" || string(event.Data) != `{"status":"working"}` {
		t.Fatalf("unexpected event: %#v", event)
	}
}

func TestNewClientRequiresAPIKey(t *testing.T) {
	if _, err := NewClient(ClientOptions{}); err == nil {
		t.Fatal("expected missing API key error")
	}
}

func TestNewMobileClientRequiresRuntimeIdentity(t *testing.T) {
	_, err := NewMobileClient(MobileClientOptions{ClientOptions: ClientOptions{APIKey: "key"}})
	if err == nil {
		t.Fatal("expected missing agent id error")
	}
	_, err = NewMobileClient(MobileClientOptions{
		ClientOptions: ClientOptions{APIKey: "key"}, AgentID: "agent-1",
	})
	if err == nil {
		t.Fatal("expected missing instance id error")
	}
}

func TestTerminalTaskStatus(t *testing.T) {
	for _, status := range []TaskStatus{TASKSTATUS_COMPLETED, TASKSTATUS_FAILED, TASKSTATUS_CANCELED, TASKSTATUS_TIMEOUT, TASKSTATUS_REJECTED, TaskStatus("cancelled")} {
		if !isTerminalTaskStatus(status) {
			t.Fatalf("status %q is not terminal", status)
		}
	}
	if isTerminalTaskStatus(TASKSTATUS_RUNNING) {
		t.Fatal("running must not be terminal")
	}
}
