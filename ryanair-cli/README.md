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

- `RYANAIR_COOKIE` — optional override (persisted cookies in `~/.ryanair/cookies.json`)
- `RYANAIR_REQUEST_DELAY` — rate limit (e.g. `2s`)


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
ryanair session chrome          # open Chrome, wait for cookies, save to ~/.ryanair/cookies.json
ryanair session sync            # sync cookies from an already-running Chrome on :9222
ryanair session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.ryanair/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `RYANAIR_COOKIE`.

## Rate limits

Use `RYANAIR_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `RYANAIR_REQUEST_DELAY`.

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
