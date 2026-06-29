#!/usr/bin/env python3
import subprocess
import sys

if len(sys.argv) < 4:
    print("usage: wave-run-with-timeout.py TIMEOUT OUT ERR -- CMD...", file=sys.stderr)
    raise SystemExit(2)

timeout = int(sys.argv[1])
out, err = sys.argv[2], sys.argv[3]
cmd = sys.argv[sys.argv.index("--") + 1 :]
try:
    with open(out, "w") as fo, open(err, "w") as fe:
        r = subprocess.run(cmd, stdout=fo, stderr=fe, timeout=timeout)
    raise SystemExit(r.returncode)
except subprocess.TimeoutExpired:
    raise SystemExit(124)
