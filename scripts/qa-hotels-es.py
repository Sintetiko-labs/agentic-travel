#!/usr/bin/env python3
"""Exhaustive live smoke tests for Spanish hotel CLIs."""
from __future__ import annotations

import json
import subprocess
import sys
from dataclasses import dataclass, field
from pathlib import Path

ROOT = Path(__file__).resolve().parent.parent
CLIS = [
    "barcelo", "riu", "catalonia", "h10", "palladium", "lopesan",
    "princess", "eurostars", "vincci", "sercotel", "silken",
]
CITIES = ["Madrid", "Barcelona", "Palma", "Valencia"]
NO_PRESENCE = {
    "lopesan": {"Madrid", "Barcelona", "Valencia"},
    "riu": {"Barcelona", "Valencia"},
    "palladium": {"Palma"},
    "vincci": {"Palma"},
    "h10": {"Valencia"},
}
PAGINATION_CITY = {"lopesan": "Gran Canaria"}
TIMEOUT = 90


@dataclass
class Result:
    cli: str
    build: str = "skip"
    help: str = "skip"
    cities: dict[str, str] = field(default_factory=dict)
    limit: str = "skip"
    page: str = "skip"
    read: str = "skip"
    doctor: str = "skip"
    notes: list[str] = field(default_factory=list)

    @property
    def overall(self) -> str:
        checks = [self.build, self.help, *self.cities.values(), self.limit, self.page, self.doctor]
        if self.read not in ("skip", "n/a"):
            checks.append(self.read)
        if any(c == "fail" for c in checks):
            return "FAIL"
        if any(c == "warn" for c in checks):
            return "WARN"
        return "PASS"


def run(cmd: list[str], cwd: Path) -> tuple[int, str, str]:
    try:
        p = subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=TIMEOUT)
        return p.returncode, p.stdout, p.stderr
    except subprocess.TimeoutExpired:
        return -1, "", "timeout"


def validate_search_json(data: dict) -> tuple[bool, str]:
    hotels = data.get("hotels")
    if not isinstance(hotels, list):
        return False, "missing hotels array"
    if not hotels:
        return False, "empty hotels array"
    for i, h in enumerate(hotels[:3]):
        if not h.get("name") or not (h.get("hotel_url") or h.get("id")):
            return False, f"hotel[{i}] missing fields"
    return True, ""


def test_cli(slug: str) -> Result:
    r = Result(cli=slug)
    cli_dir = ROOT / f"{slug}-cli"
    bin_path = cli_dir / slug
    code, out, err = run(["go", "build", "-o", slug, f"./cmd/{slug}"], cli_dir)
    if code != 0:
        r.build, r.notes = "fail", [(err or out).strip()[:200]]
        return r
    r.build = "pass"
    code, _, _ = run([str(bin_path), "help"], cli_dir)
    r.help = "pass" if code in (0, 2) else "fail"
    first_id = ""
    for city in CITIES:
        if slug == "silken" and city == "Madrid":
            r.cities[city] = "n/a"
            continue
        if slug in NO_PRESENCE and city in NO_PRESENCE[slug]:
            r.cities[city] = "n/a"
            continue
        code, out, err = run([str(bin_path), "search", "--json", city], cli_dir)
        if code != 0:
            r.cities[city] = "fail"
            r.notes.append(f"search {city}: {(out+err).strip()[:160]}")
            continue
        data = json.loads(out)
        ok, msg = validate_search_json(data)
        r.cities[city] = "pass" if ok else "fail"
        if not ok:
            r.notes.append(f"search {city}: {msg}")
        elif not first_id:
            first_id = data["hotels"][0].get("id") or data["hotels"][0].get("hotel_url", "")
    pc = PAGINATION_CITY.get(slug, "Barcelona")
    code, out, _ = run([str(bin_path), "search", "--json", "--limit", "5", pc], cli_dir)
    if code == 0:
        n = len(json.loads(out).get("hotels", []))
        r.limit = "pass" if 0 < n <= 5 else "fail"
    else:
        r.limit = "fail"
    code, out, _ = run([str(bin_path), "search", "--json", "--limit", "5", "--page", "2", pc], cli_dir)
    if code == 0:
        d = json.loads(out)
        r.page = "pass" if d.get("page") == 2 and d.get("hotels") else "warn"
    else:
        r.page = "fail"
    if first_id:
        code, out, err = run([str(bin_path), "read", "--json", first_id], cli_dir)
        if code != 0:
            r.read = "fail"
            r.notes.append(f"read: {(out+err).strip()[:120]}")
        else:
            r.read = "pass" if json.loads(out).get("name") else "warn"
    else:
        r.read = "skip"
    code, out, _ = run([str(bin_path), "session", "doctor", "--json"], cli_dir)
    r.doctor = "pass" if code == 0 and json.loads(out) else "fail"
    bin_path.unlink(missing_ok=True)
    return r


def main() -> int:
    results = [test_cli(s) for s in CLIS]
    out = [{"cli": r.cli, "overall": r.overall, "build": r.build, "help": r.help,
            "cities": r.cities, "limit": r.limit, "page": r.page, "read": r.read,
            "doctor": r.doctor, "notes": r.notes} for r in results]
    path = ROOT / "scripts" / "qa-hotels-es-results.json"
    path.write_text(json.dumps(out, indent=2))
    print(json.dumps(out, indent=2))
    return 1 if any(r.overall == "FAIL" for r in results) else 0


if __name__ == "__main__":
    sys.exit(main())
