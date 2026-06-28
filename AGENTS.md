# Guía para agentes de IA — agentic-travel

Documentación orientada a **agentes autónomos** que orquestan búsqueda de hoteles y vuelos mediante **MCP oficial (primero)** y **CLIs de este monorepo (fallback / marcas)**.

> **No oficial (CLIs).** APIs reverse-engineered. Puede romperse cuando el proveedor cambie el sitio. **Ejecutar solo en local** (IP residencial).

---

## MCP-first (loop 7)

**Regla por defecto:** intenta **MCP** para descubrimiento agregado; usa **CLI** cuando el usuario nombra una marca, cuando MCP no cubre el inventario, o cuando necesitas tarifas directas / sesión WAF.

| Intención | Herramienta primaria | Fallback |
|-----------|---------------------|----------|
| Vuelo ciudad→ciudad sin aerolínea | **Duffel MCP** | CLI LCC (`ryanair`, `vueling`, `volotea`, `binter`) |
| Hotel por ciudad / POI | **Amadeus** o **Booking.com MCP** | CLI cadena española (`melia`, `barcelo`, `nh`, …) |
| "Solo Meliá / Ryanair / …" | **CLI** (`groups.json` → slug) | MCP solo si CLI falla |
| Disponibilidad / precio miembro | **CLI** | MCP no sustituye BFF de marca |
| Reserva / orden | MCP con contrato (Duffel/Amadeus/Booking) | CLI entrega `booking_url` |

Arquitectura completa, costes y matriz de fiabilidad: **[docs/MCP_VS_CLI.md](docs/MCP_VS_CLI.md)**.

### Variables de entorno MCP (orquestador)

```bash
DUFFEL_API_KEY=...          # vuelos
AMADEUS_CLIENT_ID=...       # vuelos + hoteles GDS
AMADEUS_CLIENT_SECRET=...
BOOKING_PARTNER_ID=...      # hoteles agregador (si aplica)
```

Sin claves MCP → degradar a CLI donde exista implementación **live** o **partial**; no inventar ofertas.

### Flujo híbrido recomendado

1. Parsear destino, fechas y **marca explícita** (si hay).
2. Sin marca → MCP search (vuelos y/o hoteles).
3. Con marca o MCP vacío → `{slug} search --json` (ver `scripts/groups.json`).
4. Normalizar a tipos `travelkit` (`hotels[]`, `flights[]` — nunca `null`).
5. Detalle / availability → MCP `offer_id` si existe; si no, CLI `read` / `availability`.
6. Devolver `booking_url` de la capa que respondió.

---

## PARALLEL SEARCH PROTOCOL (loop 7)

**Objetivo:** respuestas agregadas en **<15s** en Mac Mini M-series. Guía completa: **[docs/FAST_SEARCH.md](docs/FAST_SEARCH.md)**.

### Reglas (obligatorias para agentes)

| Regla | Detalle |
|-------|---------|
| **NUNCA** CLIs secuenciales en multi-marca | Un `for slug in …` suma latencias (4×30s = 2 min). Usa orquestadores. |
| **Vuelos N marcas** | `./scripts/parallel-flights.sh --from ORIGIN --to DEST --depart DATE` |
| **Hoteles N cadenas** | `./scripts/parallel-hotels.sh --city CITY` |
| **MAD→LON vuelos + hotel** | `./scripts/wave-search-madrid-london.sh` → `wave-result.json` |
| **MCP + CLI misma ola** | Duffel MCP en **background** mientras corren parallel CLIs — no esperar MCP antes de CLIs |
| **Concurrencia** | Máx **8–10** procesos (Go orchestrator: `workers = NumCPU()`) |
| **Timeout** | **30s por fuente**; devolver **resultados parciales** |
| **1 marca nombrada** | Excepción: un solo CLI (`ryanair search --json …`) sin orquestador |

### Árbol de decisión rápido

```
¿1 marca?     → CLI único
¿N marcas?    → parallel-flights.sh / parallel-hotels.sh
¿Agregado?    → MCP (Duffel) async + parallel CLI en la misma ola
¿Confirmar?   → MCP top 3, luego CLI read solo en los 3 mejores
```

### Comandos clave

```bash
# Pre-build (una vez)
./scripts/parallel-search/build-bins.sh

# Multi-LCC
./scripts/parallel-flights.sh --from MAD --to STN --depart 2026-07-05

# Multi-cadena hoteles
./scripts/parallel-hotels.sh --city London --limit 10

# Híbrido: Madrid London flights + hotel July (<15s target)
WAVE_DEPART=2026-07-05 WAVE_HOTELS=London ./scripts/wave-search-madrid-london.sh

# MCP vuelos (background, misma ola que parallel CLIs)
MCP_FROM=MAD MCP_TO=STN MCP_DEPART=2026-07-05 ./scripts/mcp-travel-search.sh &
```

Variables: `DUFFEL_ACCESS_TOKEN`, `WAVE_DEPART`, `WAVE_OUT`, `AGENTIC_TRAVEL_BINS`.

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
- MCP: documentado en [docs/MCP_VS_CLI.md](docs/MCP_VS_CLI.md); adaptadores `travelkit/mcp/` — roadmap loop 7
- Prioridad loop 7: router híbrido MCP+CLI, registry `preferred_tool` en `groups.json`
