# TUI CLI

Unofficial, agent-friendly CLI for [TUI](https://www.tui.co.uk).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o tui ./cmd/tui
```

## Commands

```bash
tui search [--json] [--brand BRAND] --from MAD --to BCN --depart 2026-07-01
tui read [--json] [--brand BRAND] <id|url>
tui brands
```

## Environment

- `TUI_COOKIE` — optional browser cookie when blocked
- `TUI_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

- TUI Airways
- TUI fly

Use `--brand` to select a sub-brand.

## Status

Category: **airline**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
