#!/usr/bin/env node
/**
 * Gondola MCP hotel search via Streamable HTTP.
 * Usage: node mcp/call-gondola-search.mjs --city London --check-in 2026-07-05 --check-out 2026-07-08
 */
import { spawnSync } from 'node:child_process';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '..');
const GONDOLA_URL = process.env.GONDOLA_MCP_URL ?? 'https://mcp.gondola.ai/mcp';

function usage() {
  console.error(`Usage: node mcp/call-gondola-search.mjs --city CITY --check-in DATE --check-out DATE [--guests N]`);
  process.exit(2);
}

function parseArgs(argv) {
  const out = { guests: 2 };
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--city') out.location = argv[++i];
    else if (a === '--check-in') out.check_in = argv[++i];
    else if (a === '--check-out') out.check_out = argv[++i];
    else if (a === '--guests') out.guests = Number(argv[++i]);
    else if (a === '--help' || a === '-h') usage();
  }
  if (!out.location || !out.check_in || !out.check_out) usage();
  return out;
}

const args = parseArgs(process.argv.slice(2));
const toolArgs = JSON.stringify(args);
const helper = path.join(ROOT, 'mcp/call-mcp-http.mjs');

const result = spawnSync(
  process.execPath,
  [helper, '--url', GONDOLA_URL, '--tool', 'search_hotels', '--args', toolArgs],
  { stdio: 'inherit', env: process.env },
);

process.exit(result.status ?? 1);
