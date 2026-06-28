# RIU CLI

Unofficial, agent-friendly CLI for [RIU](https://www.riu.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o riu ./cmd/riu
```

## Commands

```bash
riu search [--json] [--limit N] <destination>
riu read [--json] <id|url>
riu availability [--json] --check-in DATE --check-out DATE <hotel-id>
riu brands
```

## Environment

- `RIU_COOKIE` — optional browser cookie when blocked
- `RIU_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

| Feature | Status |
|---------|--------|
| `search` | **live** — destination ng-state (`/es/hotels/europa/espana/{city}`) |
| `read` / `availability` | implemented |
| Rate limit | `RIU_REQUEST_DELAY` (~2s) |

```bash
riu search --json Madrid
```
