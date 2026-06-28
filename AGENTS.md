# Guía para agentes de IA — agentic-travel

Documentación orientada a **agentes autónomos** que orquestan búsqueda de hoteles y vuelos mediante **MCP oficial (primero)**, **browser MCP (partials WAF)** y **CLIs de este monorepo (marcas / cadenas regionales)**.

> **No oficial (CLIs).** APIs reverse-engineered. Puede romperse cuando el proveedor cambie el sitio. **Ejecutar solo en local** (IP residencial).

---

## MCP-first (loop 7)

**Regla por defecto:** intenta **MCP agregado** para descubrimiento; usa **CLI** cuando el usuario nombra una marca española, cuando MCP no cubre el inventario, o cuando necesitas tarifas directas / sesión WAF.

| Intención | Herramienta primaria | Fallback |
|-----------|---------------------|----------|
| Vuelo ciudad→ciudad sin aerolínea | **Kiwi MCP** (`search-flight`) o **Duffel MCP** | CLI LCC (`ryanair`, `vueling`, `volotea`, `binter`) |
| Hotel por ciudad / cadenas globales | **Gondola MCP** (`search_hotels`) o **Duffel** `search_stays` | CLI cadena española (`melia`, `barcelo`, `nh`, …) |
| "Solo Meliá / Ryanair / …" | **CLI** (`groups.json` → slug) | MCP solo si CLI falla |
| Partial CLI + doctor `blocked` (Akamai) | **Browser MCP** (`cursor-ide-browser`) | `{slug} session chrome --wait` en Mac con Chrome |
| Disponibilidad / precio miembro | **CLI** | Gondola puede cubrir Marriott/Hilton/Accor; MCP no sustituye BFF Meliá/NH |
| Reserva / orden | MCP con contrato (Kiwi/Duffel/Gondola links) | CLI entrega `booking_url` |

Arquitectura completa: **[docs/MCP_VS_CLI.md](docs/MCP_VS_CLI.md)** · inventario: **[docs/MCP_TRAVEL_INVENTORY.md](docs/MCP_TRAVEL_INVENTORY.md)**.

### Variables de entorno MCP (orquestador)

```bash
DUFFEL_ACCESS_TOKEN=...     # Duffel MCP — vuelos + stays (test: duffel_test_…)
AMADEUS_CLIENT_ID=...       # opcional — Amadeus community MCP
AMADEUS_CLIENT_SECRET=...
EXPEDIA_API_KEY=...         # opcional — Expedia recommendations MCP
BOOKING_PARTNER_ID=...      # opcional — Booking.com partner (cuando haya MCP público)
```

Kiwi (`https://mcp.kiwi.com`) y Gondola (`https://mcp.gondola.ai/mcp`) **no requieren clave**.

Configuración Cursor: [`.cursor/mcp.json`](.cursor/mcp.json) · guía: [docs/MCP_SETUP.md](docs/MCP_SETUP.md).

Sin claves Duffel → usar **Kiwi** para vuelos agregados y **Gondola** para hoteles de cadena; degradar a CLI donde exista implementación **live** o **partial**; no inventar ofertas.

### Browser MCP — partials con WAF

Cuando `{slug} session doctor --json` devuelve **`blocked`** y el slug está en [`bridge/browser-mcp/registry.json`](bridge/browser-mcp/registry.json), usa el **browser MCP bridge** (`cursor-ide-browser` en [`.cursor/mcp.json.example`](.cursor/mcp.json.example); fallback `@playwright/mcp`) antes de rendirte:

```
browser_navigate → browser_wait_for (Akamai) → browser_network_requests / browser_snapshot
```

Por cada marca, `registry.json` define `start_url`, `network_filter` / `dom_fallback` y el adaptador `bridge/browser-mcp/adapters/{slug}.mjs` (mismo shape `travelkit` que `{slug} search --json`). Playbooks: [bridge/browser-mcp/prompts/](bridge/browser-mcp/prompts/) (p. ej. [madrid-london.md](bridge/browser-mcp/prompts/madrid-london.md)). Arquitectura: [docs/BROWSER_MCP_BRIDGE.md](docs/BROWSER_MCP_BRIDGE.md).

**Cadenas españolas regionales** (Barceló, H10, Hotusa, Palladium, Catalonia, …): seguir en **CLI** — agregadores MCP no cubren inventario miembro.

### Flujo híbrido recomendado (MCP-first)

1. Parsear destino, fechas y **marca explícita** (si hay).
2. **Sin marca** → MCP agregado: vuelos **Kiwi** (o Duffel si hay token); hoteles **Gondola** (cadenas) o Duffel `search_stays` (genérico).
3. **Con marca española** (Meliá, Barceló, NH, Ryanair, …) → **CLI** `{slug} search --json` (ver `scripts/groups.json`).
4. Si doctor = `blocked` y slug ∈ `registry.json` → **browser MCP bridge** (playbook + adapter) o `session chrome`.
5. MCP vacío en ruta agregada → CLI LCC o cadena según `scripts/groups.json`.
6. Normalizar a tipos `travelkit` (`hotels[]`, `flights[]` — nunca `null`).
7. Detalle / availability → MCP `offer_id` / booking link si existe; si no, CLI `read` / `availability`.
8. Devolver `booking_url` de la capa que respondió.

Smoke MAD→London: `./scripts/mcp-smoke-madrid-london.sh`

---

## Propósito del monorepo (capa CLI)

`agentic-travel` agrupa CLIs Go con un **contrato común**:

- Binario estático, subcomandos por operación
- Salida estructurada con `--json` en **stdout**; errores en **stderr**
- Código de salida `1` en error
- Rate limiting configurable por env var `{PREFIX}_REQUEST_DELAY`
- Marcas del mismo grupo comparten CLI con flag `--brand`

---

## Comandos por categoría

### Hoteles

```bash
<binario> search [--json] [--brand BRAND] [--limit N] <destino...>
<binario> read [--json] [--brand BRAND] <id|url>
<binario> availability [--json] --check-in DATE --check-out DATE [--guests N] [--rooms N] <hotel-id>
<binario> brands
```

### Aerolíneas

```bash
<binario> search [--json] [--brand BRAND] --from ORIGIN --to DEST --depart DATE [--return DATE]
<binario> read [--json] [--brand BRAND] <id|url>
<binario> brands
```

---

## Instalación / build

```bash
cd <slug>-cli
go build -o <slug> ./cmd/<slug>
```

Ejemplo:

```bash
cd melia-cli && go build -o melia ./cmd/melia
cd ryanair-cli && go build -o ryanair ./cmd/ryanair
```

---

## Rate limiting

Todas las CLIs respetan `{PREFIX}_REQUEST_DELAY`:

```bash
export MELIA_REQUEST_DELAY=2s
export RYANAIR_REQUEST_DELAY=1s
```

Cookie opcional cuando hay anti-bot: `{PREFIX}_COOKIE`.

Sesión Akamai/Incapsula:

```bash
melia session chrome --wait --timeout 3m
melia session doctor --json
```

---

## Agrupación de marcas

Consulta `scripts/groups.json` o el subcomando `brands` del CLI correspondiente.

Ejemplo: `melia brands` lista Meliá, Paradisus, INNSiDE, etc.

**Cadenas españolas sin MCP fiable** (priorizar CLI): Meliá, Barceló, NH, Iberostar, H10, Hotusa, Palladium, Catalonia, Eurostars, RIU, Vincci, Silken, Sercotel, Paradores, …

**LCCs** (priorizar CLI): Ryanair, Vueling, Volotea, Binter, easyJet (partial).

---

## Librería compartida

[`travelkit/`](travelkit/) — HTTP client (uTLS), cookies, rate limit, tipos JSON normalizados (contrato común MCP ↔ CLI).

```go
replace github.com/fbelchi/travelkit => ../travelkit
```

---

## Verificación

```bash
./scripts/verify-clis.sh
```

---

## Estado del proyecto (loop 6 → 7)

- **18 live** + **7 partial** CLIs prioritarios (ver README)
- MCP integrados: **Kiwi**, **Gondola**, **Duffel**, **cursor-ide-browser** — [docs/MCP_SETUP.md](docs/MCP_SETUP.md)
- Arquitectura híbrida: [docs/MCP_VS_CLI.md](docs/MCP_VS_CLI.md) · inventario loop 7: [docs/MCP_TRAVEL_INVENTORY.md](docs/MCP_TRAVEL_INVENTORY.md)
- Prioridad loop 7: router híbrido MCP+CLI+browser, registry `preferred_tool` en `groups.json`
