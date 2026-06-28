# NH Hotel Group CLI

Unofficial, agent-friendly CLI for [NH Hotel Group](https://www.nh-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o nh ./cmd/nh
```

## Commands

```bash
nh search [--json] [--limit N] [--brand BRAND] <destination>
nh read [--json] [--brand BRAND] <id|url>
nh availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
nh brands
```

## Environment

- `NH_COOKIE` — optional override (persisted cookies in `~/.nh/cookies.json`)
- `NH_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the NH Hotel Group booking API:

- NH Hotel Group
- NH Hotels
- NH Collection
- nhow

Use `--brand` to select a sub-brand when searching.


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
nh session chrome          # open Chrome, wait for cookies, save to ~/.nh/cookies.json
nh session sync            # sync cookies from an already-running Chrome on :9222
nh session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.nh/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `NH_COOKIE`.

## Rate limits

Use `NH_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `NH_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `/nh/es/api/v1/hotels/search`; needs `NH_COOKIE` |
| `read` / `availability` | implemented |
| Rate limit | `NH_REQUEST_DELAY` (~2s) |
