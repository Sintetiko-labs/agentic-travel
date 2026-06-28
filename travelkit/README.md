# travelkit

Shared Go library for **agentic-travel** CLIs: HTTP transport with Chrome-like TLS (uTLS), cookie jar, rate limiting, and normalized JSON types for hotels and flights.

## Packages

| Package | Purpose |
|---------|---------|
| `base` | HTTP client skeleton (`GetJSON`, `FetchHTML`, headers) |
| `types` | `HotelSearchResult`, `FlightSearchResult`, etc. |
| `transport` | uTLS Chrome fingerprint |
| `cookies` | Cookie jar helpers |
| `ratelimit` | `{PREFIX}_REQUEST_DELAY` pacer |
| `httputil` | Small I/O helpers |

## Usage

```go
import tkbase "github.com/fbelchi/travelkit/base"

c := tkbase.New("https://www.melia.com", "melia")
```

Each CLI uses `replace github.com/fbelchi/travelkit => ../travelkit` in its `go.mod`.
