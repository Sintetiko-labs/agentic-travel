# Enter Air CLI

Unofficial, agent-friendly CLI for [Enter Air](https://www.enterair.pl).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o enterair ./cmd/enterair
```

## Commands

```bash
enterair search [--json] --from MAD --to BCN --depart 2026-07-01
enterair read [--json] <id|url>
enterair brands
```

## Environment

- `ENTERAIR_COOKIE` — optional browser cookie when blocked
- `ENTERAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
