# Go Webhook Delivery Pipeline

[![CI](https://github.com/FelipeMadson/go-webhook-delivery-pipeline/actions/workflows/ci.yml/badge.svg)](https://github.com/FelipeMadson/go-webhook-delivery-pipeline/actions)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Release](https://img.shields.io/badge/release-v1.0.0-green.svg)](https://github.com/FelipeMadson/go-webhook-delivery-pipeline/releases)

High-throughput, resilient webhook delivery engine written in **Go 1.22+**. Author: [Felipe Madison](https://github.com/FelipeMadson). Features bounded worker queues, constant-time HMAC-SHA256 signature verification, and backoff retry semantics.

## Architecture

```text
┌────────────────────────────────────────────────────────────────────────┐
│                        INGRESS DISPATCHER                              │
│   • Bounded Memory Channel Queue (Backpressure Prevention)             │
│   • Cryptographic HMAC-SHA256 Signature Generation                     │
└───────────────────────────────────┬────────────────────────────────────┘
                                    │
┌───────────────────────────────────▼────────────────────────────────────┐
│                        CONCURRENT WORKER POOL                          │
│   • Non-blocking Goroutine Dispatchers                                 │
│   • Exponential Backoff & Jitter Scheduling                            │
└────────────────────────────────────────────────────────────────────────┘
```

## Build & Test

```bash
go build ./...
go test -v ./...
```

## Author

Authored by **[Felipe Madison](https://github.com/FelipeMadson)**.
