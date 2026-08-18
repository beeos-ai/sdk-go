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
