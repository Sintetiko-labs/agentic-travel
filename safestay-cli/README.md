# Safestay CLI

Unofficial, agent-friendly CLI for [Safestay](https://www.safestay.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o safestay ./cmd/safestay
```

## Commands

```bash
safestay search [--json] [--limit N] <destination>
safestay read [--json] <id|url>
safestay availability [--json] --check-in DATE --check-out DATE <hotel-id>
safestay brands
```

## Environment

- `SAFESTAY_COOKIE` — optional browser cookie when blocked
- `SAFESTAY_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
