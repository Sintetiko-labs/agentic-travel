# Air Europa CLI

Unofficial, agent-friendly CLI for [Air Europa](https://www.aireuropa.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o aireuropa ./cmd/aireuropa
```

## Commands

```bash
aireuropa search [--json] --from MAD --to BCN --depart 2026-07-01
aireuropa read [--json] <id|url>
aireuropa brands
```

## Environment

- `AIREUROPA_COOKIE` — optional override (persisted cookies in `~/.aireuropa/cookies.json`)
- `AIREUROPA_REQUEST_DELAY` — rate limit (e.g. `2s`)


## Session chrome

Capture Akamai/WAF cookies from Chrome (headed browser required):

```bash
aireuropa session chrome --wait --timeout 3m
aireuropa session doctor --json
```

Manual Chrome launch (if not using `--replace`):

```bash
/Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome \
  --remote-debugging-port=9222 \
  --user-data-dir=$HOME/.aireuropa/chrome-profile \
  https://example.com
```

Cookies load automatically on `search` / `read` / `availability`. Override with `AIREUROPA_COOKIE`.

## Rate limits

Use `AIREUROPA_REQUEST_DELAY=60s` for airlines (~1 req/min). Hotels: `2s` default via `AIREUROPA_REQUEST_DELAY`.

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `dapi.aireuropa.com/api/v1/flights/search` POST; needs `aireuropa session chrome --wait` |
| `read` | implemented |
| Rate limit | `AIREUROPA_REQUEST_DELAY` (~2s) |
