# Mac smoke integration report (loop-7)

**Generated:** 2026-06-28T19:05Z (UTC)  
**Integration branch:** `loop-7/smoke-integration`  
**Main SHA at fetch:** `5110754346ad51d8e6224bb11afba1db542693fc`  
**Collector:** loop-7 smoke integration agent

## Summary

After `git fetch --all`, no `docs/SMOKE_MAC_*.md` artifacts were found on `main`, `origin/main`, or any `loop-7/*` remote branch. No `loop-7/qa-fix-*` branches exist to merge (only historical `loop-6/qa-fix-*`).

Live `GOMAXPROCS=1 ./scripts/verify-clis.sh` was **deferred** while other agents held concurrent `verify-clis` / `go build` locks on this workspace; see pass rate below.

## Agents in flight

| Process / lane | Working directory / target | Status (at report time) |
|----------------|----------------------------|-------------------------|
| `verify-clis.sh` (×2) | `agentic-travel` | Running — full CLI matrix |
| `verify-clis.sh` + tee | `/tmp/final-verify.txt` | Running |
| `GOMAXPROCS=2 verify-clis` | `/tmp/verify-coordinator-main.log` | Running |
| `go build` + search | `ryanair-cli` | Running — MAD→STN smoke |
| `go build` + search | `travelodge-cli` | Running — London search |
| Per-CLI help probe | `/tmp/vf511b-*` temp bins | Running — partial matrix |
| **This integration agent** | `docs/SMOKE_MAC_REPORT.md` | Publishing placeholder report; will refresh when per-agent `SMOKE_MAC_*.md` land |

## Source inventory (`docs/SMOKE_MAC_*.md`)

| Branch | Files found |
|--------|-------------|
| `main` / `origin/main` | *(none)* |
| `origin/loop-7/*` | *(no remote loop-7 branches yet except this integration push)* |
| Local workspace | *(none)* |

When smoke agents push, re-run integration to replace rows below from:

- `docs/SMOKE_MAC_<lane>.md`
- `docs/SMOKE_MAC_<cli>.md`

## Master table

| CLI | category | status | command | notes |
|-----|----------|--------|---------|-------|
| *(pending)* | — | — | — | No `SMOKE_MAC_*.md` sources merged yet; table will be populated on next integration pass |

## Merge actions

| Action | Result |
|--------|--------|
| Merge `loop-7/qa-fix-*` → `main` | **Skipped** — no matching branches on remote |
| Push fixes to `main` | **Skipped** — nothing merged |

## Live CLI verify (`scripts/verify-clis.sh`)

| Field | Value |
|-------|-------|
| **Command** | `GOMAXPROCS=1 ./scripts/verify-clis.sh` |
| **Status** | Deferred / in progress (concurrent agent runs) |
| **Pass rate** | *See integration agent return payload* |
| **Results file** | `scripts/verify-results.txt` (may reflect in-flight partial run) |

## Next integration pass

1. `git fetch --all` and collect all `docs/SMOKE_MAC_*.md` from `loop-7/*`.
2. Merge any `loop-7/qa-fix-*` with green smoke.
3. Run `GOMAXPROCS=1 ./scripts/verify-clis.sh` when workspace idle.
4. If fixes merged, fast-forward `main`; else keep report on `loop-7/smoke-integration`.
