# Developer setup

## Primary clone

Use **`/Users/fbelchi/github/agentic-travel`** as the main working copy.

```bash
cd /Users/fbelchi/github/agentic-travel
git fetch origin && git checkout main && git pull --ff-only
```

## Mac CLI ergonomics

Run Go binaries from **Terminal.app** (not the Cursor agent shell).

```bash
chmod +x scripts/mac-build-cli.sh
./scripts/mac-build-cli.sh travelodge search --json London --limit 3
```

## Network: residential IP only

CLIs must run from your **Mac home/office Wi‑Fi** with a **residential** public IP — not VPN, cloud agent, or datacenter egress.

- **Do not** set `HTTP_PROXY`, `HTTPS_PROXY`, or `ALL_PROXY`. `travelkit/network` and `travelkit/transport` ignore proxy env vars; you'll get a stderr warning if they are set.
- Use `--no-proxy` before any subcommand to clear proxy variables for that run, e.g. `travelodge --no-proxy search London`.
- Headed Chrome session capture (`{slug} session chrome --wait`) must use the same network.

### Verify egress

```bash
./scripts/mac-build-cli.sh travelodge network doctor
curl -s https://ifconfig.me   # manual check from the same Mac
```

`network doctor` probes `ifconfig.me`, flags datacenter/hosting IPs via ip-api, and exits non-zero on datacenter egress.

## Worktrees

Only one worktree may have `main` checked out. Do not force-push.
