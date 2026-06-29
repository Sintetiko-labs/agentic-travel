# Fast parallel search (MAD → London)

One command fans out Duffel MCP (when `DUFFEL_ACCESS_TOKEN` is set), Kiwi + Gondola HTTP MCP, and cached `ryanair` / `vueling` / `travelodge` / `hilton` CLIs into `wave-result.json` with per-source timings in `sources[]`.

```bash
./scripts/wave-search-madrid-london.sh
```
