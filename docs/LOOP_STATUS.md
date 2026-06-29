# Loop status (parallel workstreams)

Last updated: 2026-06-29 (airlines-parent-apis merge onto origin/main)

## Main branch

| Field | Value |
|-------|-------|
| **SHA** | `adf8fcd` (`adf8fcd8c353d678cec7941e2b85ead485fbe8a2`) |
| **Tip** | Loop 7: **12 airline parent APIs** (incl. new `iberia`), hotels parent/stub wiring, loop-6 QA |
| **Remote** | https://github.com/Sintetiko-labs/agentic-travel (`origin/main`) |
| **Worktree** | `/private/tmp/agentic-travel-merge` (integration); primary `/Users/fbelchi/github/agentic-travel` |

## CLI implementation counts (all *-cli with search.go)

| Status | Count | Notes |
|--------|------:|-------|
| **live / partial (non-stub)** | **79** | Real Search implementation |
| **stub** | **116** | `search not yet implemented` scaffold |
| **Total CLIs** | **194** | Session subcommands on all scaffolds |

## Loop 7 parent APIs

### Hotels

| Slug | Status | Notes |
|------|--------|-------|
| accor | **live** | JSON-LD + destination HTML (travelkit/parse/accor.go) |
| ihg | **live** | Property JSON + JSON-LD (travelkit/parse/ihg.go) |
| mamashelter, 25hours | **live** | Bespoke homepage parsers |

### Airlines (parent-group search — `loop-7/airlines-parent-apis`, commit `8ecab06`)

| Slug | Status | Notes |
|------|--------|-------|
| lufthansagroup | **live** | LH lowestfares + Eurowings API |
| airfranceklm | **live** | Air France / KLM availability |
| britishairways | **live** | BA search API |
| turkish | **live** | Turkish Airlines group |
| norwegian | **live** | Norwegian API |
| jet2 | **live** | Jet2 search |
| tui | **live** | TUI fly group |
| emirates | **live** | Emirates availability |
| qatar | **live** | Qatar Airways |
| etihad | **live** | Etihad |
| wizzair | **live** | `be.wizzair.com` search API |
| iberia | **live** | Iberia + Iberia Express + Air Nostrum (`--brand`) |

Mac smoke: `./scripts/smoke-mac-airlines-parent-apis.sh` (MAD→LHR, 90s timeout per CLI).

## Merged in consolidation batch (c9c29da → d18c8cf)

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

Merged this pass: `loop-7/airlines-parent-apis` (`8ecab06` → merge `adf8fcd`). Earlier: stub-wire-parent, hotels-parent-apis-1, wave-mcp-parallel, orchestrator, MCP stack.

## Priority README table

Documented **live** priority slugs: **35**; **partial**: **4** (marriott, easyjet, aireuropa, iberiaexpress). See root README.md.

## Verify

`./scripts/verify-clis.sh` on clean main after pull.

## Next actions

1. Wire remaining 117 stubs via scripts/wire-stub-to-parent.py + scripts/groups.json.
2. Primary clone: `git fetch origin && git checkout main && git pull origin main`.
