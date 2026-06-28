# Barceló CLI

Unofficial, agent-friendly CLI for [Barceló](https://www.barcelo.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o barcelo ./cmd/barcelo
```

## Commands

```bash
barcelo search [--json] [--limit N] [--brand BRAND] <destination>
barcelo read [--json] [--brand BRAND] <id|url>
barcelo availability [--json] [--brand BRAND] --check-in DATE --check-out DATE <hotel-id>
barcelo brands
```

## Environment

- `BARCELO_COOKIE` — optional override (persisted cookies in `~/.barcelo/cookies.json`)
- `BARCELO_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Sub-brands

This CLI covers multiple brands sharing the Barceló booking API:

- Barceló Hotel Group
- Barceló Hotels & Resorts
- Royal Hideaway
- Occidental Hotels & Resorts
- Allegro Hotels

Use `--brand` to select a sub-brand when searching.


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
barcelo session chrome          # open Chrome, wait for cookies, save to ~/.barcelo/cookies.json
barcelo session sync            # sync cookies from an already-running Chrome on :9222
barcelo session chrome --no-wait  # immediate capture
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.barcelo/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `BARCELO_COOKIE`.

## Rate limits

Use `BARCELO_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `BARCELO_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **live** — JSON-LD from `/es/hoteles` |
| `read` / `availability` | implemented |
| Rate limit | `BARCELO_REQUEST_DELAY` (~2s) |

```bash
barcelo search --json Madrid
```
