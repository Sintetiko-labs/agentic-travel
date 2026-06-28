# TOC Hostels CLI

Unofficial, agent-friendly CLI for [TOC Hostels](https://www.tochostels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o toc ./cmd/toc
```

## Commands

```bash
toc search [--json] [--limit N] <destination>
toc read [--json] <id|url>
toc availability [--json] --check-in DATE --check-out DATE <hotel-id>
toc brands
```

## Environment

- `TOC_COOKIE` — optional browser cookie when blocked
- `TOC_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
