# Browser MCP bridge (prototype)

Agent-side reliability layer for WAF-blocked partial CLIs. See [docs/BROWSER_MCP_BRIDGE.md](../../docs/BROWSER_MCP_BRIDGE.md).

## Quick start

1. Enable **cursor-ide-browser** in Cursor (built-in) or add Playwright MCP to `.cursor/mcp.json`.
2. Follow [prompts/madrid-london.md](./prompts/madrid-london.md) in an agent chat.
3. Save captured JSON/HTML from `browser_network_requests` or snapshot.
4. Normalize:

```bash
cat /tmp/melia-bff.json | node adapters/melia.mjs --stdin --query Madrid
node adapters/marriott.mjs --html /tmp/marriott.html --query London
cat /tmp/ej.json | node adapters/easyjet.mjs --stdin --origin MAD --dest LTN --depart 2026-07-15
```

Output matches `travelkit/types` with `source` suffixed `-browser-mcp`.
