#!/usr/bin/env node
/**
 * Kiwi.com MCP flight search via Streamable HTTP.
 * Usage: node mcp/call-kiwi-search.mjs --from MAD --to STN --depart 2026-07-05
 */
import { spawnSync } from 'node:child_process';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '..');
const KIWI_URL = process.env.KIWI_MCP_URL ?? 'https://mcp.kiwi.com';

function usage() {
  console.error(`Usage: node mcp/call-kiwi-search.mjs --from ORIGIN --to DEST --depart YYYY-MM-DD [--return YYYY-MM-DD] [--adults N]`);
  process.exit(2);
}

function parseArgs(argv) {
  const out = { adults: 1, cabin: 'economy' };
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--from') out.origin = argv[++i];
    else if (a === '--to') out.destination = argv[++i];
    else if (a === '--depart') out.departureDate = argv[++i];
    else if (a === '--return') out.returnDate = argv[++i];
    else if (a === '--adults') out.adults = Number(argv[++i]);
    else if (a === '--help' || a === '-h') usage();
  }
  if (!out.origin || !out.destination || !out.departureDate) usage();
  return out;
}

const args = parseArgs(process.argv.slice(2));
const toolArgs = JSON.stringify(args);
const helper = path.join(ROOT, 'mcp/call-mcp-http.mjs');

const result = spawnSync(
  process.execPath,
  [helper, '--url', KIWI_URL, '--tool', 'search-flight', '--args', toolArgs],
  { stdio: 'inherit', env: process.env },
);

process.exit(result.status ?? 1);
