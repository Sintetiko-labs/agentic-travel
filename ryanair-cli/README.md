# Ryanair CLI

Unofficial, agent-friendly CLI for [Ryanair](https://www.ryanair.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ryanair ./cmd/ryanair
```

## Commands

```bash
ryanair search [--json] --from MAD --to BCN --depart 2026-07-01
ryanair read [--json] <id|url>
ryanair brands
```

## Environment

- `RYANAIR_COOKIE` — optional browser cookie when blocked
- `RYANAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

| Feature | Status |
|---------|--------|
| `search` | **live** — `farfnd` fare calendar (no cookie); `booking/v4/availability` when session present |
| `read` | **stable** — resolves via search |
| Rate limit | ~1 req/min recommended; set `RYANAIR_REQUEST_DELAY=2s` |

Example:

```bash
ryanair search --json --from BCN --to PMI --depart 2026-07-15
# MAD→BCN may return empty (no Ryanair route) — use routes Ryanair operates
```

Session: export `RYANAIR_COOKIE` after browser search, or set `RYANAIR_CLIENT_VERSION` on 409 errors.
