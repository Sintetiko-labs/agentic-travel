#!/usr/bin/env node
/**
 * Call a remote MCP server over Streamable HTTP (Kiwi, Gondola).
 * Falls back to legacy SSE transport when Streamable HTTP is unavailable.
 *
 * Usage:
 *   node mcp/call-mcp-http.mjs --url https://mcp.kiwi.com --tool search-flight --args '{"origin":"MAD",...}'
 */
import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StreamableHTTPClientTransport } from '@modelcontextprotocol/sdk/client/streamableHttp.js';
import { SSEClientTransport } from '@modelcontextprotocol/sdk/client/sse.js';

function usage() {
  console.error(`Usage: node mcp/call-mcp-http.mjs --url URL --tool NAME --args JSON`);
  process.exit(2);
}

function parseArgs(argv) {
  const out = { args: {} };
  for (let i = 0; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--url') out.url = argv[++i];
    else if (a === '--tool') out.tool = argv[++i];
    else if (a === '--args') out.args = JSON.parse(argv[++i]);
    else if (a === '--help' || a === '-h') usage();
  }
  if (!out.url || !out.tool) usage();
  return out;
}

async function connectClient(url) {
  const client = new Client({ name: 'agentic-travel-mcp-http', version: '1.0.0' }, { capabilities: {} });
  const parsed = new URL(url);

  try {
    const transport = new StreamableHTTPClientTransport(parsed);
    await client.connect(transport);
    return { client, transport: 'streamable-http' };
  } catch (err) {
    const transport = new SSEClientTransport(parsed);
    await client.connect(transport);
    return { client, transport: 'sse' };
  }
}

function extractText(result) {
  return (result.content ?? [])
    .filter((c) => c.type === 'text')
    .map((c) => c.text)
    .join('\n');
}

const { url, tool, args } = parseArgs(process.argv.slice(2));
let client;

try {
  ({ client } = await connectClient(url));
  const result = await client.callTool({ name: tool, arguments: args });
  const text = extractText(result);
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
} catch (err) {
  console.error(`MCP HTTP call failed (${url} ${tool}): ${err.message}`);
  process.exit(1);
} finally {
  if (client) await client.close().catch(() => {});
}
