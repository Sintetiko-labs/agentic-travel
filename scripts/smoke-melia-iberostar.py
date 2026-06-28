#!/usr/bin/env python3
"""CDP smoke harness for melia and iberostar."""
import subprocess, sys
from pathlib import Path
ROOT = Path(__file__).resolve().parent.parent

def main():
    for slug, args in [("melia", ["search", "--json", "Madrid", "--limit", "3"]), ("iberostar", ["search", "--json", "Madrid", "--limit", "3"])]:
        d = ROOT / f"{slug}-cli"
        b = d / slug
        subprocess.check_call(["go", "build", "-o", slug, f"./cmd/{slug}"], cwd=d)
        subprocess.run([str(b), "session", "chrome", "--replace", "--wait"], cwd=d)
        subprocess.run([str(b), *args], cwd=d)
    return 0
if __name__ == "__main__": sys.exit(main())
