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

- `AIREUROPA_COOKIE` — optional browser cookie when blocked
- `AIREUROPA_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

| Feature | Status |
|---------|--------|
| `search` | **partial** — `/ae/api/v1/flights/search`; needs `AIREUROPA_COOKIE` |
| `read` | implemented |
| Rate limit | `AIREUROPA_REQUEST_DELAY` (~2s) |
