# Madrid → London — Browser MCP agent playbook

**Route:** Madrid origin hotels + easyJet to London + Marriott London hotels  
**Date:** 2026-07-15  
**Prereq:** `cursor-ide-browser` or Playwright MCP in Cursor

## A. easyJet MAD → LTN

1. `browser_navigate(url="https://www.easyjet.com/es")`
2. `browser_wait_for(time=5)` → `browser_snapshot()` → fill MAD / LTN / 2026-07-15 → search
3. `browser_network_requests()` — filter `ejavailability/api/v5/availability/query`
4. `node bridge/browser-mcp/adapters/easyjet.mjs --file /tmp/ej.json --origin MAD --dest LTN --depart 2026-07-15`

## B. Meliá Madrid

1. `browser_navigate(url="https://www.melia.com/es/hoteles")` → wait 8s → search **Madrid**
2. Capture POST `.../services/search/hotels/v2/search`
3. `node bridge/browser-mcp/adapters/melia.mjs --file /tmp/melia-bff.json --query Madrid`

## C. NH Madrid

1. `browser_navigate(url="https://www.nh-hotels.com/es/hoteles/espana/madrid")`
2. Capture GET `.../nh/es/api/v1/hotels/search`
3. `node bridge/browser-mcp/adapters/nh.mjs --file /tmp/nh.json --query Madrid`

## D. Marriott London

1. Deep-link `findHotels.mi` with London + dates 07/15–07/16/2026
2. Save HTML if no JSON in network log
3. `node bridge/browser-mcp/adapters/marriott.mjs --html /tmp/marriott.html --query London`

## Failures

| Symptom | Fix |
|---------|-----|
| Tool `browser_navigate` missing | Enable cursor-ide-browser or add `@playwright/mcp` |
| Empty network log | Longer wait; retry navigate |
| 0 parsed hotels | Screenshot + console messages |
