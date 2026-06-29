# Loop status (parallel workstreams)

Last updated: 2026-06-29 (full consolidation onto origin/main)

## Main branch

| Field | Value |
|-------|-------|
| **SHA** | `9f446b0` (`9f446b079c7f2a5101e780659ca50a2f32b7d3d7`) |
| **Tip** | Loop 7: 8 international hotel parent APIs (73 brands), stub wiring, loop-6 QA merges |
| **Remote** | https://github.com/Sintetiko-labs/agentic-travel (`origin/main`) |
| **Worktree** | `/private/tmp/agentic-travel-merge` (integration); primary `/Users/fbelchi/github/agentic-travel` |

## CLI implementation counts (all *-cli with search.go)

| Status | Count | Notes |
|--------|------:|-------|
| **live / partial (non-stub)** | **77** | Real Search implementation |
| **stub** | **118** | `search not yet implemented` scaffold |
| **Total CLIs** | **194** | Session subcommands on all scaffolds |

## Loop 7 parent APIs

| Slug | Status | Notes |
|------|--------|-------|
| accor | **live** | 11 brands — JSON-LD + destination HTML (travelkit/parse/accor.go) |
| ihg | **live** | 9 brands — property JSON + JSON-LD |
| hyatt | **live** | 11 brands — shared travelkit hotel helpers |
| marriott | **live** | 20 brands — Bonvoy search + `--brand` filter |
| hilton | **live** | 8 brands |
| wyndham | **live** | 5 brands |
| bestwestern | **live** | 4 brands |
| radisson | **live** | 5 brands |
| lufthansagroup | **live** | LH lowestfares + Eurowings API |
| mamashelter, 25hours | **live** | Bespoke homepage parsers (accor group; not exec-wired) |

## Merged in consolidation batch (c9c29da → d18c8cf)

| Branch | Status |
|--------|--------|
| loop-7/hotels-parent-apis-1 | merged |
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

1. Wire remaining 118 stubs via scripts/wire-stub-to-parent.py + scripts/groups.json.
2. Primary clone: `git fetch origin && git checkout main && git pull origin main`.
