#!/usr/bin/env node
import { readFileSync } from 'node:fs';

function parseArgs(argv) {
  const opts = { query: '', page: 1, pageSize: 24, html: null };
  for (let i = 2; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--query') opts.query = argv[++i] ?? '';
    else if (a === '--page') opts.page = Number(argv[++i]);
    else if (a === '--page-size') opts.pageSize = Number(argv[++i]);
    else if (a === '--html') opts.html = argv[++i];
  }
  return opts;
}

function extractJsonLd(html) {
  const re = /<script[^>]*type="application\/ld\+json"[^>]*>([\s\S]*?)<\/script>/gi;
  const hotels = [];
  let m;
  while ((m = re.exec(html)) !== null) {
    try {
      const data = JSON.parse(m[1]);
      const items = Array.isArray(data) ? data : [data];
      for (const item of items) {
        if (item['@type'] === 'Hotel' || item['@type'] === 'LodgingBusiness') {
          hotels.push({
            id: item.identifier ?? item.url ?? item.name ?? '',
            name: item.name ?? '',
            city: item.address?.addressLocality ?? '',
            country: item.address?.addressCountry ?? '',
            hotel_url: item.url ?? '',
          });
        }
      }
    } catch { /* skip */ }
  }
  return hotels;
}

function extractListingLinks(html, base = 'https://www.marriott.com') {
  const re = /href="(\/[^"]*\/overview\/[^"]+)"/gi;
  const nameRe = /data-property-name="([^"]+)"/gi;
  const names = [...html.matchAll(nameRe)].map((x) => x[1]);
  const links = [...html.matchAll(re)].map((x) => x[1]);
  const out = [];
  const seen = new Set();
  for (let i = 0; i < links.length; i++) {
    const path = links[i];
    if (seen.has(path)) continue;
    seen.add(path);
    out.push({
      id: path,
      name: names[i] ?? path.split('/').filter(Boolean).pop() ?? 'Hotel',
      hotel_url: `${base}${path}`,
    });
  }
  return out;
}

const opts = parseArgs(process.argv);
if (!opts.html) {
  console.error('usage: node marriott.mjs --html page.html --query London');
  process.exit(1);
}
const html = readFileSync(opts.html, 'utf8');
let rows = extractJsonLd(html);
if (rows.length === 0) rows = extractListingLinks(html);
const total = rows.length;
const start = (opts.page - 1) * opts.pageSize;
const slice = rows.slice(start, start + opts.pageSize).map((h) => ({
  id: String(h.id),
  name: h.name,
  brand: 'Marriott',
  city: h.city ?? opts.query,
  country: h.country ?? '',
  price: '',
  currency: 'GBP',
  hotel_url: h.hotel_url,
  image_url: '',
}));
process.stdout.write(JSON.stringify({
  query: opts.query,
  total,
  page: opts.page,
  page_size: opts.pageSize,
  has_next_page: total > opts.page * opts.pageSize,
  hotels: slice,
  brand: 'Marriott',
  source: 'findHotels-browser-mcp',
}, null, 2) + '\n');
