# VP Hoteles CLI

Unofficial, agent-friendly CLI for [VP Hoteles](https://www.vp-hoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o vp ./cmd/vp
```

## Commands

```bash
vp search [--json] [--limit N] <destination>
vp read [--json] <id|url>
vp availability [--json] --check-in DATE --check-out DATE <hotel-id>
vp brands
```

## Environment

- `VP_COOKIE` — optional browser cookie when blocked
- `VP_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
