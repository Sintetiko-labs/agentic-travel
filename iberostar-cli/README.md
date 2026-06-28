# Iberostar CLI

Unofficial, agent-friendly CLI for [Iberostar](https://www.iberostar.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o iberostar ./cmd/iberostar
```

## Commands

```bash
iberostar search [--json] [--limit N] [--brand BRAND] <destination>
iberostar read [--json] [--brand BRAND] <id|url>
iberostar availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
iberostar brands
```

## Environment

- `IBEROSTAR_COOKIE` — optional override (persisted cookies in `~/.iberostar/cookies.json`)
- `IBEROSTAR_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Iberostar booking API:

- Iberostar
- Iberostar Selection
- Iberostar Grand

Use `--brand` to select a sub-brand when searching.


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
iberostar session chrome          # open Chrome, wait for cookies, save to ~/.iberostar/cookies.json
iberostar session sync            # sync cookies from an already-running Chrome on :9222
iberostar session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.iberostar/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `IBEROSTAR_COOKIE`.

## Rate limits

Use `IBEROSTAR_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `IBEROSTAR_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — GraphQL `/api/graphql`; needs `IBEROSTAR_COOKIE` |
| `read` / `availability` | implemented |
| Rate limit | `IBEROSTAR_REQUEST_DELAY` (~2s) |
