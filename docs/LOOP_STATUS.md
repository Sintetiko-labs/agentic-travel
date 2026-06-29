# Loop status (parallel workstreams)

Last updated: 2026-06-29 (full consolidation onto origin/main)

## Main branch

| Field | Value |
|-------|-------|
| **SHA** | `af95548` (`af9554827d93226095696f2982bf631daeb0803b`) |
| **Tip** | Loop 7 consolidation: parent APIs (accor/ihg/lufthansagroup), stub wiring, loop-6 QA merges |
| **Remote** | https://github.com/Sintetiko-labs/agentic-travel (`origin/main`) |
| **Worktree** | `/private/tmp/agentic-travel-merge` (integration); primary `/Users/fbelchi/github/agentic-travel` |

## CLI implementation counts (all *-cli with search.go)

| Status | Count | Notes |
|--------|------:|-------|
| **live / partial (non-stub)** | **77** | Real Search implementation |
| **stub** | **117** | `search not yet implemented` scaffold |
| **Total CLIs** | **194** | Session subcommands on all scaffolds |

## Loop 7 parent APIs

| Slug | Status | Notes |
|------|--------|-------|
| accor | **live** | JSON-LD + destination HTML (travelkit/parse/accor.go) |
| ihg | **live** | Property JSON + JSON-LD (travelkit/parse/ihg.go) |
| lufthansagroup | **live** | LH lowestfares + Eurowings API |
| mamashelter, 25hours | **live** | Bespoke homepage parsers |

## Merged in consolidation batch (c9c29da → af95548)

| Branch | Status |
|--------|--------|
| loop-7/hotels-es-stubs-live | merged |
| loop-7/recovery-status | merged |
| loop-7/mcp-setup | merged |
| loop-7/learnings-ecommartinez | merged |
| loop-7/waf-integration | merged |
| loop-6/qa-integration | merged |
| loop-6/hotels-akamai | merged |
| loop-6/qa-fix-* (12 hotel QA branches) | merged |

Earlier on main: stub-wire-parent, hotels-parent-apis-1, airlines-parent-apis, wave-mcp-parallel, orchestrator, MCP stack.

## Priority README table

Documented **live** priority slugs: **35**; **partial**: **4** (marriott, easyjet, aireuropa, iberiaexpress). See root README.md.

## Verify

`./scripts/verify-clis.sh` on clean main after pull.

## Next actions

1. Wire remaining 117 stubs via scripts/wire-stub-to-parent.py + scripts/groups.json.
2. Primary clone: `git fetch origin && git checkout main && git pull origin main`.
