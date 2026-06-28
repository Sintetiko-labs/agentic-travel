# Iberia Express CLI

Unofficial, agent-friendly CLI for [Iberia Express](https://www.iberiaexpress.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberiaexpress ./cmd/iberiaexpress
```

## Commands

```bash
iberiaexpress search [--json] --from MAD --to BCN --depart 2026-07-01
iberiaexpress read [--json] <id|url>
iberiaexpress brands
```

## Environment

- `IBERIAEXPRESS_COOKIE` — optional override (persisted cookies in `~/.iberiaexpress/cookies.json`)
- `IBERIAEXPRESS_REQUEST_DELAY` — rate limit (e.g. `2s`)


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
iberiaexpress session chrome          # open Chrome, wait for cookies, save to ~/.iberiaexpress/cookies.json
iberiaexpress session sync            # sync cookies from an already-running Chrome on :9222
iberiaexpress session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.iberiaexpress/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `IBERIAEXPRESS_COOKIE`.

## Rate limits

Use `IBERIAEXPRESS_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `IBERIAEXPRESS_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `/api/availability/v1/flights`; needs `IBERIAEXPRESS_COOKIE` |
| `read` | implemented |
| Rate limit | `IBERIAEXPRESS_REQUEST_DELAY` (~2s) |
