# ITA Airways CLI

Unofficial, agent-friendly CLI for [ITA Airways](https://www.ita-airways.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ita ./cmd/ita
```

## Commands

```bash
ita search [--json] --from MAD --to BCN --depart 2026-07-01
ita read [--json] <id|url>
ita brands
```

## Environment

- `ITA_COOKIE` — optional browser cookie when blocked
- `ITA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
