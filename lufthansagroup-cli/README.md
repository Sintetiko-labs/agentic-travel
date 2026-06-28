# Lufthansa Group CLI

Unofficial, agent-friendly CLI for [Lufthansa Group](https://www.lufthansa.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o lufthansagroup ./cmd/lufthansagroup
```

## Commands

```bash
lufthansagroup search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
lufthansagroup read [--json] [--brand BRAND] <id|url>
lufthansagroup brands
```

## Environment

- `LUFTHANSAGROUP_COOKIE` — optional browser cookie when blocked
- `LUFTHANSAGROUP_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- Lufthansa
- Lufthansa City Airlines
- Discover Airlines
- Swiss
- Austrian Airlines
- Brussels Airlines
- Eurowings

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
