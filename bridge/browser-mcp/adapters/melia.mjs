#!/usr/bin/env node
import { readFileSync } from 'node:fs';

function parseArgs(argv) {
  const opts = { query: '', page: 1, pageSize: 24, stdin: false, file: null };
  for (let i = 2; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--stdin') opts.stdin = true;
    else if (a === '--query') opts.query = argv[++i] ?? '';
    else if (a === '--page') opts.page = Number(argv[++i]);
    else if (a === '--page-size') opts.pageSize = Number(argv[++i]);
    else if (a === '--file') opts.file = argv[++i];
  }
  return opts;
}

function loadInput(opts) {
  if (opts.stdin) return readFileSync(0, 'utf8');
  if (opts.file) return readFileSync(opts.file, 'utf8');
  throw new Error('provide --stdin or --file');
}

function mapHotel(row, base) {
  const code = row.code ?? row.id ?? '';
  const url = row.url ?? row.hotelUrl ?? '';
  const absUrl = url.startsWith('http') ? url : `${base}${url}`;
  const price = row.minPrice ?? row.price;
  return {
    id: String(code),
    name: row.name ?? '',
    brand: row.brand ?? 'Meliá',
    city: row.city ?? '',
    country: row.country ?? '',
    stars: row.category ?? row.stars ?? 0,
    price: price != null ? String(price) : '',
    currency: row.currency ?? 'EUR',
    hotel_url: absUrl,
    image_url: row.image ? (row.image.startsWith('http') ? row.image : `${base}${row.image}`) : '',
  };
}

function toResult(raw, query, page, pageSize) {
  const base = 'https://www.melia.com';
  const hotels = (raw.hotels ?? raw.data?.hotels ?? raw.data ?? []).map((row) => mapHotel(row, base));
  const total = raw.total ?? hotels.length;
  const start = (page - 1) * pageSize;
  return {
    query,
    total,
    page,
    page_size: pageSize,
    has_next_page: total > page * pageSize,
    hotels: hotels.slice(start, start + pageSize),
    brand: 'Meliá',
    source: 'melia-bff-browser-mcp',
  };
}

const opts = parseArgs(process.argv);
const raw = JSON.parse(loadInput(opts));
process.stdout.write(JSON.stringify(toResult(raw, opts.query, opts.page, opts.pageSize), null, 2) + '\n');
