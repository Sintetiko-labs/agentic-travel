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

- `RIU_COOKIE` — optional override (persisted cookies in `~/.riu/cookies.json`)
- `RIU_REQUEST_DELAY` — rate limit (e.g. `2s`)


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
riu session chrome          # open Chrome, wait for cookies, save to ~/.riu/cookies.json
riu session sync            # sync cookies from an already-running Chrome on :9222
riu session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.riu/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `RIU_COOKIE`.

## Rate limits

Use `RIU_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `RIU_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **live** — destination ng-state (`/es/hotels/europa/espana/{city}`) |
| `read` / `availability` | implemented |
| Rate limit | `RIU_REQUEST_DELAY` (~2s) |

```bash
riu search --json Madrid
```
