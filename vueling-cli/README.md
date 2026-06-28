# Vueling CLI

Unofficial, agent-friendly CLI for [Vueling](https://www.vueling.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o vueling ./cmd/vueling
```

## Commands

```bash
vueling search [--json] --from MAD --to BCN --depart 2026-07-01
vueling read [--json] <id|url>
vueling brands
```

## Environment

- `VUELING_COOKIE` — optional override (persisted cookies in `~/.vueling/cookies.json`)
- `VUELING_REQUEST_DELAY` — rate limit (e.g. `2s`)


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
vueling session chrome          # open Chrome, wait for cookies, save to ~/.vueling/cookies.json
vueling session sync            # sync cookies from an already-running Chrome on :9222
vueling session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.vueling/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `VUELING_COOKIE`.

## Rate limits

Use `VUELING_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `VUELING_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `www.vueling.com/bit/v2` 404 (SPA); fallback `tickets.vueling.com/bit/v2` (Skysales). Needs `vueling session chrome` on tickets host |
| `read` | implemented |
| Rate limit | `VUELING_REQUEST_DELAY` (~60s recommended) |
