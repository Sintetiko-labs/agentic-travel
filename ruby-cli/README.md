# Ruby Hotels CLI

Unofficial, agent-friendly CLI for [Ruby Hotels](https://www.ruby-hotels.com).

> **Not official.** Reverse-engineered endpoints. Run locally only. Respect rate limits.

## Build

```bash
go build -o ruby ./cmd/ruby
```

## Commands

```bash
ruby search [--json] [--limit N] <destination>
ruby read [--json] <id|url>
ruby availability [--json] --check-in DATE --check-out DATE <hotel-id>
ruby brands
```

## Environment

- `RUBY_COOKIE` — optional browser cookie when blocked
- `RUBY_REQUEST_DELAY` — rate limit (e.g. `2s`)

## Status

Category: **hotel**

Search: **scaffold** — TODO implement endpoint in `internal/client/search.go`
