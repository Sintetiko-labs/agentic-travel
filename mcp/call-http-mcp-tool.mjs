#!/usr/bin/env node
import { Client } from '@modelcontextprotocol/sdk/client/index.js';
import { StreamableHTTPClientTransport } from '@modelcontextprotocol/sdk/client/streamableHttp.js';

function usage() {
  console.error('Usage: node mcp/call-http-mcp-tool.mjs --url URL --tool NAME --args JSON');
  process.exit(2);
}

let url, tool, argsJson = '{}';
for (let i = 2; i < process.argv.length; i++) {
  const a = process.argv[i];
  if (a === '--url') url = process.argv[++i];
  else if (a === '--tool') tool = process.argv[++i];
  else if (a === '--args') argsJson = process.argv[++i];
  else if (a === '--help' || a === '-h') usage();
}
if (!url || !tool) usage();

let toolArgs;
try {
  toolArgs = JSON.parse(argsJson);
} catch (e) {
  console.error('Invalid --args JSON:', e.message);
  process.exit(2);
}

const transport = new StreamableHTTPClientTransport(new URL(url));
const client = new Client({ name: 'agentic-travel-wave', version: '1.0.0' }, { capabilities: {} });

try {
  await client.connect(transport);
  const result = await client.callTool({ name: tool, arguments: toolArgs });
  const text = (result.content ?? [])
    .filter((c) => c.type === 'text')
    .map((c) => c.text)
    .join('\n');
  let payload;
  if (text) {
    try {
      payload = JSON.parse(text);
    } catch {
      payload = { text };
    }
  } else {
    payload = result;
  }
  console.log(JSON.stringify(payload));
  if (result.isError) process.exit(1);
} finally {
  await client.close();
}
