# Ilunion CLI

Unofficial, agent-friendly CLI for [Ilunion](https://www.ilunionhotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ilunion ./cmd/ilunion
```

## Commands

```bash
ilunion search [--json] [--limit N] <destination>
ilunion read [--json] <id|url>
ilunion availability [--json] --check-in DATE --check-out DATE <hotel-id>
ilunion brands
```

## Environment

- `ILUNION_COOKIE` — optional browser cookie when blocked
- `ILUNION_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
