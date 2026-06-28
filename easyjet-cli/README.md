# easyJet CLI

Unofficial, agent-friendly CLI for [easyJet](https://www.easyjet.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o easyjet ./cmd/easyjet
```

## Commands

```bash
easyjet search [--json] --from MAD --to BCN --depart 2026-07-01
easyjet read [--json] <id|url>
easyjet brands
```

## Environment

- `EASYJET_COOKIE` — optional override (persisted cookies in `~/.easyjet/cookies.json`)
- `EASYJET_REQUEST_DELAY` — rate limit (e.g. `2s`)


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
easyjet session chrome          # open Chrome, wait for cookies, save to ~/.easyjet/cookies.json
easyjet session sync            # sync cookies from an already-running Chrome on :9222
easyjet session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.easyjet/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `EASYJET_COOKIE`.

## Rate limits

Use `EASYJET_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `EASYJET_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `ejavailability/api/v5`; needs `EASYJET_COOKIE` |
| `read` | implemented |
| Rate limit | `EASYJET_REQUEST_DELAY` (~2s) |
