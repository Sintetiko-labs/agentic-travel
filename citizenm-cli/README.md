# citizenM CLI

Unofficial, agent-friendly CLI for [citizenM](https://www.citizenm.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o citizenm ./cmd/citizenm
```

## Commands

```bash
citizenm search [--json] [--limit N] <destination>
citizenm read [--json] <id|url>
citizenm availability [--json] --check-in DATE --check-out DATE <hotel-id>
citizenm brands
```

## Environment

- `CITIZENM_COOKIE` — optional browser cookie when blocked
- `CITIZENM_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
