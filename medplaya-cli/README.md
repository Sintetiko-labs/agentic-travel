# MedPlaya CLI

Unofficial, agent-friendly CLI for [MedPlaya](https://www.medplaya.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o medplaya ./cmd/medplaya
```

## Commands

```bash
medplaya search [--json] [--limit N] <destination>
medplaya read [--json] <id|url>
medplaya availability [--json] --check-in DATE --check-out DATE <hotel-id>
medplaya brands
```

## Environment

- `MEDPLAYA_COOKIE` — optional browser cookie when blocked
- `MEDPLAYA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
