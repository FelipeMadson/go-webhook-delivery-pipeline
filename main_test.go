package main

import (
	"testing"
)

func TestSignPayload(t *testing.T) {
	pipeline := NewDeliveryPipeline("test-secret", 10, 3)
	data := []byte("hello world")
	sig := pipeline.SignPayload(data)

	if len(sig) != 64 {
		t.Fatalf("Expected 64-character SHA-256 HMAC signature, got %d", len(sig))
	}
}

func TestEnqueue(t *testing.T) {
	pipeline := NewDeliveryPipeline("test-secret", 2, 3)
	ok1 := pipeline.Enqueue(WebhookPayload{ID: "1"})
	ok2 := pipeline.Enqueue(WebhookPayload{ID: "2"})
	ok3 := pipeline.Enqueue(WebhookPayload{ID: "3"})

	if !ok1 || !ok2 {
		t.Fatalf("Expected first two payloads to be enqueued successfully")
	}
	if ok3 {
		t.Fatalf("Expected third payload to fail due to bounded capacity")
	}
	if pipeline.QueueLen() != 2 {
		t.Fatalf("Expected queue length 2, got %d", pipeline.QueueLen())
	}
}
