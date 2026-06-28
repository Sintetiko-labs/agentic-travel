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

const base = 'https://www.nh-hotels.com';

function mapRow(h) {
  return {
    id: String(h.id ?? ''),
    name: h.name ?? '',
    brand: h.brand ?? 'NH',
    city: h.city ?? '',
    country: h.country ?? '',
    stars: h.stars ?? 0,
    price: h.price > 0 ? String(h.price) : '',
    currency: h.currency ?? 'EUR',
    hotel_url: h.slug ? `${base}/es/hotel/${h.slug}` : base,
    image_url: h.image?.startsWith('http') ? h.image : h.image ? `${base}${h.image}` : '',
  };
}

function toResult(raw, query, page, pageSize) {
  const rows = raw.data?.length ? raw.data : raw.hotels ?? [];
  const hotels = rows.map(mapRow);
  const total = raw.total ?? hotels.length;
  const start = (page - 1) * pageSize;
  return {
    query,
    total,
    page,
    page_size: pageSize,
    has_next_page: total > page * pageSize,
    hotels: hotels.slice(start, start + pageSize),
    brand: 'NH',
    source: 'nh-api-browser-mcp',
  };
}

const opts = parseArgs(process.argv);
const body = opts.stdin ? readFileSync(0, 'utf8') : readFileSync(opts.file, 'utf8');
process.stdout.write(JSON.stringify(toResult(JSON.parse(body), opts.query, opts.page, opts.pageSize), null, 2) + '\n');
