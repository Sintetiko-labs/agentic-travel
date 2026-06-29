#!/usr/bin/env python3
"""Mac live smoke tests for Spanish hotel CLIs (search/read/doctor)."""
from __future__ import annotations

import json
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
CLIS = [
    "barcelo", "riu", "catalonia", "h10", "palladium", "lopesan",
    "princess", "eurostars", "vincci", "sercotel", "silken",
    "globales", "grupotel", "hipotels", "senator", "medplaya", "zenit",
    "abba", "porthotels", "ona", "belive", "evenia", "ilunion",
    "petitpalace", "paradores", "roommate", "onlyyou", "pinero",
    "hotusa", "melia", "nh", "iberostar",
]
CITY_OVERRIDES = {
    "silken": "Barcelona",
    "grupotel": "Barcelona",
    "hipotels": "Barcelona",
    "medplaya": "Barcelona",
    "evenia": "Barcelona",
    "belive": "Barcelona",
    "pinero": "Barcelona",
    "paradores": "Segovia",
}
TIMEOUT = 120


def run(cmd: list[str], cwd: Path) -> tuple[int, str, str]:
    try:
        p = subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=TIMEOUT)
        return p.returncode, p.stdout, p.stderr
    except subprocess.TimeoutExpired:
        return -1, "", "timeout"


def smoke_cli(name: str) -> dict:
    r: dict = {"cli": name, "build": "FAIL", "search": "skip", "read": "skip", "doctor": "skip", "status": "FAIL", "notes": []}
    cli_dir = ROOT / f"{name}-cli"
    bin_path = Path(f"/tmp/{name}")
    code, out, err = run(["go", "build", "-o", str(bin_path), f"./cmd/{name}"], cli_dir)
    if code != 0:
        r["notes"].append((err or out).strip()[:200])
        return r
    r["build"] = "PASS"

    city = CITY_OVERRIDES.get(name, "Madrid")
    code, out, err = run([str(bin_path), "search", "--json", city, "--limit", "3"], cli_dir)
    combined = out + err
    if code == -1:
        r["search"] = "BLOCKED"
        r["status"] = "BLOCKED"
        r["notes"].append(f"search timeout ({TIMEOUT}s)")
        return r
    if "search not yet implemented" in combined.lower():
        r["search"] = "FAIL"
        r["notes"].append("search not yet implemented")
        return r
    if code != 0:
        r["search"] = "FAIL"
        r["notes"].append(f"search exit {code}: {combined.strip()[:200]}")
        return r
    try:
        data = json.loads(out)
    except json.JSONDecodeError as e:
        r["search"] = "FAIL"
        r["notes"].append(f"invalid JSON: {e}")
        return r
    hotels = data.get("hotels") or []
    names = [h["name"] for h in hotels if h.get("name")]
    if not names:
        r["search"] = "FAIL"
        r["notes"].append("empty hotels or missing names")
        return r
    r["search"] = "PASS"
    r["city"] = city
    r["hotel_names"] = names[:3]

    hid = hotels[0].get("id") or hotels[0].get("hotel_url") or ""
    if hid:
        code, rout, rerr = run([str(bin_path), "read", "--json", str(hid)], cli_dir)
        if code == 0:
            try:
                rd = json.loads(rout)
                r["read"] = "PASS" if rd.get("name") else "FAIL"
                if not rd.get("name"):
                    r["notes"].append("read missing name")
            except json.JSONDecodeError:
                r["read"] = "FAIL"
                r["notes"].append("read invalid JSON")
        else:
            r["read"] = "FAIL"
            r["notes"].append(f"read exit {code}: {(rout + rerr).strip()[:120]}")

    code, dout, derr = run([str(bin_path), "session", "doctor", "--json"], cli_dir)
    if code == 0:
        try:
            json.loads(dout)
            r["doctor"] = "PASS"
        except json.JSONDecodeError:
            r["doctor"] = "FAIL"
            r["notes"].append("doctor invalid JSON")
    else:
        r["doctor"] = "FAIL"
        r["notes"].append(f"doctor exit {code}: {(dout + derr).strip()[:120]}")

    if r["search"] == "PASS" and r["doctor"] == "PASS" and r["read"] in ("PASS", "skip"):
        r["status"] = "PASS"
    elif r["search"] == "BLOCKED":
        r["status"] = "BLOCKED"
    else:
        r["status"] = "FAIL"
    return r


def main() -> int:
    out_path = Path(sys.argv[1]) if len(sys.argv) > 1 else Path("/tmp/smoke-hotels-es-results.json")
    results = [smoke_cli(c) for c in CLIS]
    out_path.write_text(json.dumps(results, indent=2))
    print(json.dumps(results, indent=2))
    passed = sum(1 for r in results if r["status"] == "PASS")
    failed = sum(1 for r in results if r["status"] == "FAIL")
    blocked = sum(1 for r in results if r["status"] == "BLOCKED")
    print(f"\nSummary: pass={passed} fail={failed} blocked={blocked}", file=sys.stderr)
    return 0 if failed == 0 and blocked == 0 else 1


if __name__ == "__main__":
    sys.exit(main())
