package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sync"
)

type WebhookPayload struct {
	ID        string
	TargetURL string
	Data      []byte
	Attempt   int
	Delivered bool
}

type DeliveryPipeline struct {
	secretKey  string
	workQueue  chan WebhookPayload
	maxRetries int
	wg         sync.WaitGroup
}

func NewDeliveryPipeline(secretKey string, queueCapacity int, maxRetries int) *DeliveryPipeline {
	return &DeliveryPipeline{
		secretKey:  secretKey,
		workQueue:  make(chan WebhookPayload, queueCapacity),
		maxRetries: maxRetries,
	}
}

func (p *DeliveryPipeline) SignPayload(data []byte) string {
	mac := hmac.New(sha256.New, []byte(p.secretKey))
	mac.Write(data)
	return hex.EncodeToString(mac.Sum(nil))
}

func (p *DeliveryPipeline) Enqueue(payload WebhookPayload) bool {
	select {
	case p.workQueue <- payload:
		return true
	default:
		return false
	}
}

func (p *DeliveryPipeline) QueueLen() int {
	return len(p.workQueue)
}

func main() {
	pipeline := NewDeliveryPipeline("production-secret-signing-key", 100, 3)
	payload := WebhookPayload{
		ID:        "evt-001",
		TargetURL: "https://api.partner.io/webhooks",
		Data:      []byte(`{"event":"payment.confirmed"}`),
		Attempt:   0,
	}
	signature := pipeline.SignPayload(payload.Data)
	fmt.Printf("Webhook Pipeline Initialized. Signature: %s\n", signature)
}
