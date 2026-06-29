# Recovery status (loop-7)

Generated: 2026-06-29 (lightweight audit; no build/verify-clis).

## Main (`origin/main`)

| Field | Value |
|-------|-------|
| **HEAD SHA** | `9dcdd3525b0ed57206fbb8dd154f34ccb2ed9d3d` |
| **Tip commit** | `9dcdd35` Merge remote-tracking branch 'origin/loop-7/transport-perf' |

**Note:** `main` is checked out in worktree `/private/tmp/agentic-travel-merge`, so this run used `git fetch --all` and `origin/main` instead of `git checkout main && git pull` in this clone.

## Recovery feature paths on `origin/main`

| Path | On main |
|------|---------|
| `orchestrator/` | yes |
| `scripts/parallel-search/` | yes |
| `scripts/mac-build-all.sh` | yes |
| `travelkit/network/` | yes |
| `bridge/browser-mcp/` | yes |
| `mcp/` | yes |
| `docs/FAST_SEARCH.md` | yes |

**Merged recovery features (checklist):** **yes** — all listed paths are present on `origin/main`.

## `origin/loop-7/*` not merged into `main`

Count: **5**

```
origin/loop-7/fix-marriott-waf
origin/loop-7/hotels-intl-batch-5
origin/loop-7/mcp-architecture
origin/loop-7/mcp-setup
origin/loop-7/wave-mcp-cli
```

## Manual merge (user)

From a clone where `main` is available (close conflicting worktree if needed):

```bash
cd /Users/fbelchi/github/agentic-travel   # or your clone
git fetch --all
git checkout main && git pull

for branch in \
  loop-7/fix-marriott-waf \
  loop-7/hotels-intl-batch-5 \
  loop-7/mcp-architecture \
  loop-7/mcp-setup \
  loop-7/wave-mcp-cli
do
  echo "=== merging origin/$branch ==="
  git merge --no-ff "origin/$branch" -m "Merge origin/$branch into main"
done
```

Resolve conflicts per branch; run your usual tests before pushing `main`.

## This branch

- Branch: `loop-7/recovery-status`
- Purpose: documentation-only recovery snapshot; **not** merged to `main` unless you choose to.
