# Royal Jordanian CLI

Unofficial, agent-friendly CLI for [Royal Jordanian](https://www.rj.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o royaljordanian ./cmd/royaljordanian
```

## Commands

```bash
royaljordanian search [--json] --from MAD --to BCN --depart 2026-07-01
royaljordanian read [--json] <id|url>
royaljordanian brands
```

## Environment

- `ROYALJORDANIAN_COOKIE` — optional browser cookie when blocked
- `ROYALJORDANIAN_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
