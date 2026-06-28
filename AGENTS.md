# Guía para agentes de IA — agentic-travel

Documentación orientada a **agentes autónomos** que orquestan búsqueda de hoteles y vuelos mediante **MCP oficial (primero)**, **browser MCP (partials WAF)** y **CLIs de este monorepo (marcas / cadenas regionales)**.

> **No oficial (CLIs).** APIs reverse-engineered. Puede romperse cuando el proveedor cambie el sitio. **Ejecutar solo en local** (IP residencial).

---

## MCP-first (loop 7)

**Regla por defecto:** intenta **MCP agregado** para descubrimiento; usa **CLI** cuando el usuario nombra una marca, cuando MCP no cubre el inventario, o cuando necesitas tarifas directas / sesión WAF.

| Intención | Herramienta primaria | Fallback |
|-----------|---------------------|----------|
| Vuelo ciudad→ciudad sin aerolínea | **Duffel MCP** | CLI LCC (`ryanair`, `vueling`, `volotea`, `binter`) |
| Hotel por ciudad / POI | **Amadeus** o **Booking.com MCP** | CLI cadena española (`melia`, `barcelo`, `nh`, …) |
| "Solo Meliá / Ryanair / …" | **CLI** (`groups.json` → slug) | MCP solo si CLI falla |
| Partial CLI + doctor `blocked` (Akamai) | **Browser MCP** (`cursor-ide-browser`) | `{slug} session chrome --wait` en Mac con Chrome |
| Disponibilidad / precio miembro | **CLI** | MCP no sustituye BFF de marca |
| Reserva / orden | MCP con contrato (Duffel/Amadeus/Booking) | CLI entrega `booking_url` |

Arquitectura completa: **[docs/MCP_VS_CLI.md](docs/MCP_VS_CLI.md)**.

### Variables de entorno MCP (orquestador)

```bash
DUFFEL_ACCESS_TOKEN=...     # vuelos + stays (Duffel MCP)
AMADEUS_CLIENT_ID=...       # vuelos + hoteles GDS
AMADEUS_CLIENT_SECRET=...
BOOKING_PARTNER_ID=...      # hoteles agregador (si aplica)
```

Sin claves MCP → degradar a CLI donde exista implementación **live** o **partial**; no inventar ofertas.

Configuración Cursor: [`.cursor/mcp.json`](.cursor/mcp.json) · guía: [docs/MCP_SETUP.md](docs/MCP_SETUP.md) · auditoría: [docs/MCP_LOCAL_AUDIT.md](docs/MCP_LOCAL_AUDIT.md).

### Browser MCP — partials con WAF (melia, nh, marriott)

Cuando `{slug} session doctor --json` devuelve **`blocked`** y no hay cookies en `~/.{slug}/cookies.json`, usa **`cursor-ide-browser`** (habilitado en `.cursor/mcp.json`) antes de rendirte:

```
browser_navigate → esperar (Akamai) → browser_snapshot → extraer JSON de red o DOM
```

| Marca | URL inicio | Filtro red / extracción |
|-------|------------|-------------------------|
| **melia** | `https://www.melia.com/es/hoteles` | POST `…/services/search/hotels/v2/search` |
| **nh** | `https://www.nh-hotels.com/es/hoteles/espana/madrid` | GET `…/nh/es/api/v1/hotels/search` |
| **marriott** | `findHotels.mi` con ciudad + fechas | JSON de red o DOM / JSON-LD |

Normalizar con `bridge/browser-mcp/adapters/{slug}.mjs` → mismo shape `travelkit` que `{slug} search --json`. Playbook: [bridge/browser-mcp/prompts/madrid-london.md](bridge/browser-mcp/prompts/madrid-london.md). Arquitectura: [docs/BROWSER_MCP_BRIDGE.md](docs/BROWSER_MCP_BRIDGE.md).

**Cadenas españolas regionales** (Barceló, H10, Hotusa, Palladium, Catalonia, …): seguir en **CLI** — agregadores MCP no cubren inventario miembro.

### Flujo híbrido recomendado

1. Parsear destino, fechas y **marca explícita** (si hay).
2. Sin marca → **Duffel MCP** (vuelos) y/o Amadeus/Booking (hoteles).
3. Con marca → `{slug} search --json`; si doctor = `blocked` → **browser MCP** o `session chrome`.
4. MCP vacío en ruta agregada → CLI LCC o cadena según `scripts/groups.json`.
5. Normalizar a tipos `travelkit` (`hotels[]`, `flights[]` — nunca `null`).
6. Detalle / availability → MCP `offer_id` si existe; si no, CLI `read` / `availability`.
7. Devolver `booking_url` de la capa que respondió.

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
- MCP: [docs/MCP_SETUP.md](docs/MCP_SETUP.md), [docs/MCP_LOCAL_AUDIT.md](docs/MCP_LOCAL_AUDIT.md), [docs/MCP_VS_CLI.md](docs/MCP_VS_CLI.md)
- Browser bridge: [docs/BROWSER_MCP_BRIDGE.md](docs/BROWSER_MCP_BRIDGE.md) · `.cursor/mcp.json` con `cursor-ide-browser`
- Prioridad loop 7: router híbrido MCP+CLI+browser, registry `preferred_tool` en `groups.json`
