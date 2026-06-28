# Kenya Airways CLI

Unofficial, agent-friendly CLI for [Kenya Airways](https://www.kenya-airways.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o kenyaairways ./cmd/kenyaairways
```

## Commands

```bash
kenyaairways search [--json] --from MAD --to BCN --depart 2026-07-01
kenyaairways read [--json] <id|url>
kenyaairways brands
```

## Environment

- `KENYAAIRWAYS_COOKIE` — optional browser cookie when blocked
- `KENYAAIRWAYS_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
