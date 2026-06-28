#!/usr/bin/env python3
from __future__ import annotations
import argparse, json, sys
from pathlib import Path
from typing import Any

def load_payload(path: Path) -> Any:
    text = path.read_text(encoding="utf-8").strip()
    if not text: return None
    try: return json.loads(text)
    except json.JSONDecodeError: return {"raw": text}

def classify(name: str, payload: Any) -> str:
    if name in ("travelodge", "hilton"): return "hotels"
    if name == "vueling" and isinstance(payload, dict) and payload.get("legs"): return "flights_connect"
    return "flights"

def count_totals(payload: Any, kind: str) -> int:
    if not isinstance(payload, dict): return 0
    if kind == "flights_connect": return int(payload.get("total") or 0)
    return int(payload.get("total") or len(payload.get("flights") or payload.get("hotels") or []) or 0)

def merge_vueling_legs(leg_a: Any, leg_b: Any, ms: int) -> dict[str, Any]:
    total = 0
    if isinstance(leg_a, dict): total += int(leg_a.get("total") or 0)
    if isinstance(leg_b, dict): total += int(leg_b.get("total") or 0)
    return {"mode": "connect", "legs": [{"route": "MAD-BCN", "data": leg_a}, {"route": "BCN-LGW", "data": leg_b}], "total": total, "ms": ms}

def main() -> int:
    ap = argparse.ArgumentParser()
    ap.add_argument("--out", required=True)
    ap.add_argument("--wall-ms", type=int, required=True)
    ap.add_argument("--meta", default="{}")
    ap.add_argument("sources", nargs="+")
    args = ap.parse_args()
    try: meta = json.loads(args.meta)
    except json.JSONDecodeError as exc:
        print(f"invalid --meta JSON: {exc}", file=sys.stderr); return 2
    sources_out, flights_total, hotels_total = [], 0, 0
    pending_vueling = None
    for spec in args.sources:
        parts = spec.split(":")
        if len(parts) < 3: print(f"bad source spec: {spec}", file=sys.stderr); return 2
        name, path, ms = parts[0], Path(parts[1]), int(parts[2])
        ok = parts[3] not in ("0", "false", "fail") if len(parts) >= 4 else True
        payload, err = None, None
        if path.is_file():
            try: payload = load_payload(path)
            except OSError as exc: ok, err = False, str(exc)
        else: ok, err = False, "missing output file"
        if name == "vueling_leg1":
            pending_vueling = {"leg1": payload, "ms": ms, "ok": ok}; continue
        if name == "vueling_leg2":
            if pending_vueling is None: pending_vueling = {"leg1": None, "ms": ms, "ok": ok}
            payload = merge_vueling_legs(pending_vueling.get("leg1"), payload, pending_vueling.get("ms", ms) + ms)
            name, ok = "vueling", ok and pending_vueling.get("ok", True)
            pending_vueling = None
        kind = classify(name, payload)
        total = count_totals(payload, kind)
        if kind == "hotels": hotels_total += total
        else: flights_total += total
        entry = {"name": name, "kind": kind, "ms": ms, "ok": ok, "total": total}
        if err: entry["error"] = err
        if payload is not None: entry["data"] = payload
        sources_out.append(entry)
    if pending_vueling is not None:
        sources_out.append({"name": "vueling", "kind": "flights", "ms": pending_vueling.get("ms", 0), "ok": False, "total": 0, "error": "incomplete connect legs", "data": pending_vueling.get("leg1")})
    result = {**meta, "wall_ms": args.wall_ms, "sources": sources_out, "flights_total": flights_total, "hotels_total": hotels_total}
    Path(args.out).write_text(json.dumps(result, indent=2) + "\n")
    print(json.dumps({"out": args.out, "wall_ms": args.wall_ms, "flights_total": flights_total, "hotels_total": hotels_total}))
    return 0

if __name__ == "__main__":
    raise SystemExit(main())
