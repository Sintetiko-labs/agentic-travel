# Recovery status (loop-7)

Generated: 2026-06-29 (full consolidation — local loop-6/7 branches merged to main).

## Main (origin/main)

| Field | Value |
|-------|-------|
| **HEAD SHA** | `af9554827d93226095696f2982bf631daeb0803b` |
| **Short** | `af95548` |
| **URL** | https://github.com/Sintetiko-labs/agentic-travel/commit/af9554827d93226095696f2982bf631daeb0803b |

## Recovery feature paths on main

| Path | Present |
|------|---------|
| orchestrator/ | yes |
| scripts/parallel-search/ | yes |
| scripts/wave-search-madrid-london.sh | yes |
| scripts/mac-build-all.sh | yes |
| travelkit/network/ | yes |
| travelkit/chrome/ | yes |
| bridge/browser-mcp/ | yes |
| mcp/ | yes |
| docs/FAST_SEARCH.md | yes |
| docs/MCP_SETUP.md | yes |
| docs/STUB_ELIMINATION.md | yes |

**Local-only commits on main after push:** target ZERO (origin/main matches worktree).

## Consolidation merges

Base before batch: `c9c29da`. Final: `af95548` (includes loop-7 hotels-es-stubs-live, recovery-status, mcp-setup, learnings-ecommartinez, waf-integration, loop-6 QA + hotels-akamai).

## Primary clone sync

```bash
cd /Users/fbelchi/github/agentic-travel
git fetch origin
git checkout main
git pull origin main
```

Integration worktree: `/private/tmp/agentic-travel-merge`.
