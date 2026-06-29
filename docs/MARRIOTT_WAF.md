# Marriott Akamai WAF (loop-7)

## Problem

`marriott session chrome --wait` captures `_abck` + `bm_sz`, but `marriott search` can still return HTTP 403 from `findHotels.mi`. Akamai binds session cookies to the **browser TLS fingerprint**. Programmatic requests with the same cookie header but a different TLS stack are rejected.

## Fixes on `loop-7/fix-marriott-waf`

| Layer | Change |
|-------|--------|
| **utls** | Chrome 131 utls via `travelkit/transport` (matches `DefaultUA`). |
| **Marriott auto-CDP** | When Akamai cookies are ready **and** Chrome debugging is listening, `marriott` client routes HTTP through CDP in-browser `fetch`. |
| **Search fallback** | `fetchSearchHTML` retries via CDP `fetch` on HTTP 403. |

## Workflow (Mac, residential IP)

```bash
cd marriott-cli && go build -o marriott ./cmd/marriott
./marriott session chrome --wait --timeout 3m --replace   # keep Chrome open
./marriott search --json London --limit 5
```

## Does Marriott require in-browser fetch only?

**Often yes.** Cookies minted in real Chrome may not work with utls-only requests. Reliable path: capture session in headed Chrome, keep debugging port open, let CLI use CDP `fetch`.

## Verdict matrix

| Condition | Expected |
|-----------|----------|
| No session | BLOCKED |
| Cookies + utls only | Often BLOCKED |
| Cookies + headed Chrome on CDP port | PASS (if session fresh) |
