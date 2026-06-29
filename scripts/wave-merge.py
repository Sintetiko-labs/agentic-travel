#!/usr/bin/env python3
"""Merge parallel wave search JSON fragments into wave-result.json."""
from __future__ import annotations

import argparse
import json
import sys
from datetime import datetime, timezone
from pathlib import Path
from typing import Any


def load_json(path: Path) -> Any:
    if not path.is_file():
        return None
    try:
        return json.loads(path.read_text())
    except json.JSONDecodeError:
        return {"_parse_error": True, "raw": path.read_text()[:8000]}


def main() -> int:
    ap = argparse.ArgumentParser(description="Merge wave search source files")
    ap.add_argument("--meta-dir", required=True)
    ap.add_argument("--out", required=True)
    ap.add_argument("--wall-ms", type=int, default=0)
    ap.add_argument("--query", default="{}")
    args = ap.parse_args()
    meta_dir = Path(args.meta_dir)
    try:
        query = json.loads(args.query)
    except json.JSONDecodeError:
        print("invalid --query JSON", file=sys.stderr)
        return 2

    sources: list[dict[str, Any]] = []
    flights: list[Any] = []
    hotels: list[Any] = []

    for meta_path in sorted(meta_dir.glob("*.meta.json")):
        meta = json.loads(meta_path.read_text())
        sid = meta.get("id") or meta_path.name.replace(".meta.json", "")
        body = load_json(meta_dir / f"{sid}.json")
        entry = {
            "id": sid,
            "ok": bool(meta.get("ok")),
            "ms": int(meta.get("ms") or 0),
            "skipped": bool(meta.get("skipped")),
            "exit_code": meta.get("exit_code"),
            "error": meta.get("error"),
        }
        if body is not None:
            entry["data"] = body
        sources.append(entry)
        if not entry["ok"] or entry["skipped"] or body is None or not isinstance(body, dict):
            continue
        if isinstance(body.get("flights"), list):
            flights.extend(body["flights"])
        if isinstance(body.get("hotels"), list):
            hotels.extend(body["hotels"])
        if isinstance(body.get("offers"), list) and sid in ("ryanair", "vueling"):
            flights.extend(body["offers"])

    out = {
        "generated_at": datetime.now(timezone.utc).isoformat(),
        "query": query,
        "wall_ms": args.wall_ms,
        "sources": sources,
        "flights": flights,
        "hotels": hotels,
    }
    Path(args.out).write_text(json.dumps(out, indent=2) + "\n")
    print(args.out)
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
