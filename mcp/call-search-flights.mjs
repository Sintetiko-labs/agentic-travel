#!/usr/bin/env node
/**
 * Minimal MCP client: call duffel-mcp search_flights and print JSON to stdout.
 * Usage: node mcp/call-search-flights.mjs --from MAD --to STN --depart 2026-07-05
 */
import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StdioClientTransport } from '@modelcontextprotocol/sdk/client/stdio.js';
import { fileURLToPath } from 'url';
import path from 'path';

const __dirname = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(__dirname, '..');
const DEFAULT_SERVER = path.join(ROOT, 'mcp/vendor/duffel-mcp/dist/index.js');

function usage() {
  console.error(`Usage: node mcp/call-search-flights.mjs --from ORIGIN --to DEST --depart YYYY-MM-DD [--return YYYY-MM-DD] [--adults N] [--direct]`);
  process.exit(2);
}

function parseArgs(argv) {
  const out = { adults: 1, direct: false };
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--from') out.origin = argv[++i];
    else if (a === '--to') out.destination = argv[++i];
    else if (a === '--depart') out.departure_date = argv[++i];
    else if (a === '--return') out.return_date = argv[++i];
    else if (a === '--adults') out.adults = Number(argv[++i]);
    else if (a === '--direct') out.direct = true;
    else if (a === '--help' || a === '-h') usage();
  }
  if (!out.origin || !out.destination || !out.departure_date) usage();
  return out;
}

const args = parseArgs(process.argv.slice(2));

if (!process.env.DUFFEL_ACCESS_TOKEN?.trim()) {
  console.error('DUFFEL_ACCESS_TOKEN is required (duffel_test_… from https://duffel.com dashboard, Test mode).');
  process.exit(1);
}

const serverPath = process.env.DUFFEL_MCP_SERVER ?? DEFAULT_SERVER;

const toolArgs = {
  origin: args.origin,
  destination: args.destination,
  departure_date: args.departure_date,
  adults: args.adults,
};
if (args.return_date) toolArgs.return_date = args.return_date;
if (args.direct) toolArgs.nonstop = true;

const transport = new StdioClientTransport({
  command: 'node',
  args: [serverPath],
  env: { ...process.env },
});

const client = new Client({ name: 'agentic-travel-mcp-client', version: '1.0.0' }, { capabilities: {} });

try {
  await client.connect(transport);
  const result = await client.callTool({ name: 'search_flights', arguments: toolArgs });
  const text = (result.content ?? [])
    .filter((c) => c.type === 'text')
    .map((c) => c.text)
    .join('\n');
  if (text) {
    try {
      console.log(JSON.stringify(JSON.parse(text), null, 2));
    } catch {
      console.log(text);
    }
  } else {
    console.log(JSON.stringify(result, null, 2));
  }
  if (result.isError) process.exit(1);
} finally {
  await client.close();
}
