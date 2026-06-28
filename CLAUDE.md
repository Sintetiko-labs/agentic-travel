# Claude / Cursor — agentic-travel

Quick reference para desarrollo y orquestación con Claude Code o Cursor sobre este monorepo.

---

## Quick start

| Tipo | Build | Smoke test |
|------|-------|------------|
| Meliá (hotel) | `cd melia-cli && go build -o melia ./cmd/melia` | `./melia search --json Madrid` |
| Ryanair (airline) | `cd ryanair-cli && go build -o ryanair ./cmd/ryanair` | `./ryanair search --json --from MAD --to BCN --depart 2026-07-01` |
| Marriott (hotel) | `cd marriott-cli && go build -o marriott ./cmd/marriott` | `./marriott brands` |

Requisitos: Go 1.26+

---

## Estructura

```
agentic-travel/
├── AGENTS.md
├── CLAUDE.md
├── README.md
├── travelkit/
├── melia-cli/
│   ├── cmd/melia/
│   └── internal/client/
├── ryanair-cli/
└── scripts/
    ├── scaffold-clis.py
    ├── groups.json
    └── verify-clis.sh
```

---

## Convenciones

- Un módulo Go por CLI con `replace` a `travelkit`
- Hoteles: `search`, `read`, `availability`
- Aerolíneas: `search`, `read`
- Flag `--brand` cuando el CLI agrupa sub-marcas
- Config futura: `~/.<binario>/config.toml`

---

## Regenerar scaffolds

```bash
python3 scripts/scaffold-clis.py
```

---

## Verificación batch

```bash
chmod +x scripts/verify-clis.sh
./scripts/verify-clis.sh
```
