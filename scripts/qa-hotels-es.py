#!/usr/bin/env python3
"""Live smoke tests for Spanish hotel CLIs. See docs/QA_HOTELS_ES.md."""
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
    p = subprocess.run(cmd, cwd=cwd, capture_output=True, text=True, timeout=90)
    return p.returncode, p.stdout, p.stderr


def test_cli(slug: str) -> Result:
    r = Result(cli=slug)
    d = ROOT / f"{slug}-cli"
    b = d / slug
    code, out, err = run(["go", "build", "-o", slug, f"./cmd/{slug}"], d)
    if code != 0:
        return Result(cli=slug, build="fail", notes=[(err or out)[:200]])
    r.build = "pass"
    r.help = "pass" if run([str(b), "help"], d)[0] in (0, 2) else "fail"
    first = ""
    for city in CITIES:
        if slug == "silken" and city == "Madrid":
            r.cities[city] = "n/a"
            continue
        if slug in NO_PRESENCE and city in NO_PRESENCE[slug]:
            r.cities[city] = "n/a"
            continue
        c, o, e = run([str(b), "search", "--json", city], d)
        if c != 0:
            r.cities[city] = "fail"
            r.notes.append(f"search {city}: {(o+e)[:160]}")
            continue
        data = json.loads(o)
        hotels = data.get("hotels") or []
        if not hotels or not all(h.get("name") and (h.get("hotel_url") or h.get("id")) for h in hotels[:3]):
            r.cities[city] = "fail"
        else:
            r.cities[city] = "pass"
            if not first:
                first = hotels[0].get("id") or hotels[0].get("hotel_url", "")
    pc = PAGINATION_CITY.get(slug, "Barcelona")
    c, o, _ = run([str(b), "search", "--json", "--limit", "5", pc], d)
    if c == 0:
        n = len(json.loads(o).get("hotels", []))
        r.limit = "pass" if 0 < n <= 5 else "fail"
    else:
        r.limit = "fail"
    c, o, _ = run([str(b), "search", "--json", "--limit", "5", "--page", "2", pc], d)
    if c == 0:
        d2 = json.loads(o)
        r.page = "pass" if d2.get("page") == 2 and d2.get("hotels") else "warn"
    else:
        r.page = "fail"
    if first:
        c, o, e = run([str(b), "read", "--json", first], d)
        r.read = "pass" if c == 0 and json.loads(o).get("name") else "fail"
        if r.read == "fail":
            r.notes.append(f"read: {(o+e)[:120]}")
    else:
        r.read = "skip"
    c, o, _ = run([str(b), "session", "doctor", "--json"], d)
    r.doctor = "pass" if c == 0 else "fail"
    b.unlink(missing_ok=True)
    return r


def main() -> int:
    results = []
    for s in CLIS:
        r = test_cli(s)
        results.append({
            "cli": r.cli, "overall": r.overall, "build": r.build, "help": r.help,
            "cities": r.cities, "limit": r.limit, "page": r.page, "read": r.read,
            "doctor": r.doctor, "notes": r.notes,
        })
    path = ROOT / "scripts" / "qa-hotels-es-results.json"
    path.write_text(json.dumps(results, indent=2))
    print(json.dumps(results, indent=2))
    return 1 if any(x["overall"] == "FAIL" for x in results) else 0


if __name__ == "__main__":
    sys.exit(main())
