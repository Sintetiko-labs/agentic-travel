# Guía para agentes de IA — agentic-travel

Documentación orientada a **agentes autónomos** que orquestan búsqueda de hoteles y vuelos mediante los CLIs de este monorepo.

> **No oficial.** APIs reverse-engineered. Puede romperse cuando el proveedor cambie el sitio. **Ejecutar solo en local** (IP residencial).

---

## Propósito del monorepo

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

---

## Agrupación de marcas

Consulta `scripts/groups.json` o el subcomando `brands` del CLI correspondiente.

Ejemplo: `melia brands` lista Meliá, Paradisus, INNSiDE, etc.

---

## Librería compartida

[`travelkit/`](travelkit/) — HTTP client (uTLS), cookies, rate limit, tipos JSON normalizados.

```go
replace github.com/fbelchi/travelkit => ../travelkit
```

---

## Verificación

```bash
./scripts/verify-clis.sh
```

---

## Estado del proyecto (iteración 1)

- Estructura y scaffolds completos
- Endpoints reales: **TODO** en `internal/client/search.go` y `read.go`
- Prioridad loop 2: Meliá, Barceló, Ryanair, Vueling, Iberia Express
