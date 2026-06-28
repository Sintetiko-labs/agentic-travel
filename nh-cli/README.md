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

Capture Akamai cookies from headed Chrome (`_abck` + `bm_sz` required):

```bash
nh session chrome --wait --timeout 3m   # browse nh-hotels.com until doctor passes
nh session sync
nh session doctor --json                # GET probe with locale=es
```

Manual Chrome launch:

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.nh/chrome-profile \
  https://www.nh-hotels.com/es/hoteles
```

Cookies load automatically on `search` / `read` / `availability`. Override with `NH_COOKIE`.

## Rate limits

Use `NH_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `NH_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — REST `/nh/es/api/v1/hotels/search?locale=es`; needs `NH_COOKIE` (Akamai `_abck`+`bm_sz`) |
| `read` / `availability` | implemented |
| Rate limit | `NH_REQUEST_DELAY` (~2s) |
