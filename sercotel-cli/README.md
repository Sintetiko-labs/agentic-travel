# Sercotel CLI

Unofficial, agent-friendly CLI for [Sercotel](https://www.sercotelhoteles.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o sercotel ./cmd/sercotel
```

## Commands

```bash
sercotel search [--json] [--limit N] <destination>
sercotel read [--json] <id|url>
sercotel availability [--json] --check-in DATE --check-out DATE <hotel-id>
sercotel brands
```

## Environment

- `SERCOTEL_COOKIE` — optional browser cookie when blocked
- `SERCOTEL_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **live** — Magnolia data in Next.js RSC HTML (`internal/client/search.go`, source: `rsc-json`)
