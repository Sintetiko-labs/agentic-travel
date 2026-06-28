# Meliá CLI

Unofficial, agent-friendly CLI for [Meliá](https://www.melia.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o melia ./cmd/melia
```

## Commands

```bash
melia search [--json] [--limit N] [--brand BRAND] <destination>
melia read [--json] [--brand BRAND] <id|url>
melia availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
melia brands
```

## Environment

- `MELIA_COOKIE` — optional override (persisted cookies in `~/.melia/cookies.json`)
- `MELIA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Meliá booking API:

- Meliá Hotels International
- Meliá
- Gran Meliá
- ME by Meliá
- The Meliá Collection
- Paradisus
- INNSiDE by Meliá
- Sol by Meliá
- ZEL

Use `--brand` to select a sub-brand when searching.


## Session chrome

Capture browser cookies from headed Chrome (required for live BFF search; directory fallback works with site cookies):

```bash
melia session chrome --wait --timeout 3m   # open Chrome on hotel directory, wait for cookies
melia session sync                         # sync from Chrome already on :9222
melia session doctor --json                # POST probe to BFF + cookie validation
```

`--wait` blocks until site cookies are captured (Meliá does not always set Akamai `_abck`/`bm_sz`). Browse the hotel directory if the wait times out.

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.melia/chrome-profile \
  https://www.melia.com/es/hoteles
```

Cookies load automatically on `search` / `read` / `availability`. Override with `MELIA_COOKIE`.

## Rate limits

Use `MELIA_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `MELIA_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — BFF `/services/search/hotels/v2/search` with directory fallback `/es/hoteles`; needs `MELIA_COOKIE` for live BFF |
| `read` / `availability` | implemented (cookie for live) |
| Rate limit | `MELIA_REQUEST_DELAY` (~2s) |
