# MCP local audit — Cursor workspace inventory

**Date:** 2026-06-28  
**Scope:** MCP servers available to the user in Cursor; travel relevance for `agentic-travel`.  
**Prior audit:** agent `99d34ecc` (read-only pass; this doc completes the write-up).

---

## Executive summary

| Finding | Detail |
|---------|--------|
| **Travel MCP locally** | **None enabled** — no Duffel, Amadeus, Kiwi, or hotel aggregator MCP in active Cursor/Codex configs |
| **Best indirect candidate** | **`cursor-ide-browser`** — real Chromium for Akamai/Incapsula partial CLIs (`melia`, `nh`, `marriott`, …) |
| **agentic-travel config** | Had `.cursor/mcp.json.example` (Duffel only); **no** committed `.cursor/mcp.json` until loop 7 |
| **Reference workspaces** | `suapea`, `firstmate`, `agentic-commerce` run **`cursor-ide-browser`** via Cursor project cache — not via repo `mcp.json` |

---

## Enabled MCP servers (parent workspace `Users-fbelchi-github`)

Path: `/Users/fbelchi/.cursor/projects/Users-fbelchi-github/mcps/`

| Server ID | Travel? | Tools (cached descriptors) |
|-----------|---------|------------------------------|
| `cursor-app-control` | No | `move_agent_to_root`, `move_agent_to_cloned_root`, `create_project`, `open_automation`, `rename_chat` |
| `plugin-context7-plugin-context7` | No | `resolve-library-id`, `query-docs` |
| `user-higgsfield` | No | `mcp_auth` only (OAuth; AI image/video at `https://mcp.higgsfield.ai/mcp`) |
| `plugin-stripe-stripe` | No | `mcp_auth` only (OAuth for Stripe MCP) |

**Not enabled in parent workspace:** `cursor-ide-browser`.

---

## `cursor-ide-browser` (available in other Cursor projects)

Present in project caches for:

- `Users-fbelchi-github-suapea`
- `Users-fbelchi-github-firstmate`
- `Users-fbelchi-github-agentic-commerce`
- `Users-fbelchi-github-backlog-agents`

**16 tools** (from `…/mcps/cursor-ide-browser/tools/`):

`browser_navigate`, `browser_tabs`, `browser_lock`, `browser_snapshot`, `browser_take_screenshot`, `browser_click`, `browser_type`, `browser_fill`, `browser_select_option`, `browser_press_key`, `browser_scroll`, `browser_drag`, `browser_highlight`, `browser_get_bounding_box`, `browser_mouse_click_xy`, `browser_cdp`

### Travel relevance

Not a travel API. **Strongest indirect fit** for `agentic-travel`:

- Navigate brand sites with real TLS + JS (Akamai `_abck`, Incapsula)
- Capture BFF JSON via `browser_cdp` (`Network.enable` + response bodies) or DOM via `browser_snapshot`
- Complements `{slug} session chrome` and [`bridge/browser-mcp/`](../../bridge/browser-mcp/) adapters

**Limitations:** cookie/storage CDP commands denied; iframe content inaccessible; agent-host residential IP still required.

---

## Config files searched

| Location | Result |
|----------|--------|
| `/Users/fbelchi/github/agentic-travel/.cursor/mcp.json` | **Not found** (pre loop-7) |
| `/Users/fbelchi/github/agentic-travel/.cursor/mcp.json.example` | Duffel vendored server only |
| `~/.cursor/mcp.json` | `higgsfield` remote URL only |
| `~/.codex/config.toml` | `MCP_DOCKER`, `atlassian`, `playwright` — **no travel** |
| `suapea/mcp.json` | Playwright, MongoDB, Stripe, S3, Redis, memory — **not travel** |
| `firstmate/`, `agentic-commerce/` | No repo-level `.cursor/mcp.json`; browser MCP enabled at workspace level |

`agentic-travel` today: Go CLIs + `travelkit/session` + `scripts/smoke-mac-cdp.py` (raw CDP `:9222`), not MCP.

---

## Travel MCP candidates from local setup

| Candidate | Source | Role for agentic-travel |
|-----------|--------|-------------------------|
| **cursor-ide-browser** | Cursor built-in (enable per workspace) | **Best fit** — WAF brands, browser MCP bridge |
| **playwright** | `~/.codex/config.toml`, `suapea/mcp.json` | Fallback if built-in browser MCP missing ([issue #3878](https://github.com/cursor/cursor/issues/3878)) |
| **duffel** (vendored) | `mcp/install.sh` + `.cursor/mcp.json.example` | Aggregate flights/stays when `DUFFEL_ACCESS_TOKEN` set |
| **MCP_DOCKER** | Codex | Unknown catalog until gateway inspected |

**Not travel candidates:** higgsfield (media), stripe (payments), context7 (library docs), cursor-app-control (IDE), atlassian (Jira/Confluence).

**No dedicated travel MCP** (GDS, Amadeus, Skyscanner, Booking.com) was found in local Cursor or Codex configs at audit time.

---

## Comparison: reference workspaces

| Workspace | `cursor-ide-browser` in mcps/ | Repo `mcp.json` |
|-----------|------------------------------|-----------------|
| suapea | Yes | `playwright` + infra MCPs at repo root |
| firstmate | Yes | None |
| agentic-commerce | Yes | None (see `docs/MCP-INTEGRATION.md` for provider yaml) |
| agentic-travel (before loop-7) | No | Example Duffel only |

Pattern: **browser MCP is workspace-scoped** (Cursor project cache). Project `.cursor/mcp.json` opts in built-ins and adds CLI servers (Duffel, Playwright fallback).

---

## Recommendations (implemented loop 7)

1. **Commit `.cursor/mcp.json`** in `agentic-travel` with `cursor-ide-browser` + Playwright fallback — see [MCP_SETUP.md](./MCP_SETUP.md).
2. **Add Duffel/Amadeus** entries when API keys are available (`./mcp/install.sh`, env vars).
3. **Route partial CLIs** through [BROWSER_MCP_BRIDGE.md](./BROWSER_MCP_BRIDGE.md) when `session doctor` reports `blocked`.
4. **Keep CLIs** for Spanish regional chains and LCCs — MCP does not replace brand BFF depth.

---

## Related

- [MCP_SETUP.md](./MCP_SETUP.md) — enable servers in Cursor
- [MCP_VS_CLI.md](./MCP_VS_CLI.md) — hybrid routing
- [BROWSER_MCP_BRIDGE.md](./BROWSER_MCP_BRIDGE.md) — melia / nh / marriott session pattern
- [MCP_TRAVEL_INVENTORY.md](./MCP_TRAVEL_INVENTORY.md) — external travel MCP catalog
