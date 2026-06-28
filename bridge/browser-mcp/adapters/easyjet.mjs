#!/usr/bin/env node
import { readFileSync } from 'node:fs';

function parseArgs(argv) {
  const opts = { origin: '', dest: '', depart: '', ret: '', page: 1, pageSize: 24, stdin: false, file: null };
  for (let i = 2; i < argv.length; i++) {
    const a = argv[i];
    if (a === '--stdin') opts.stdin = true;
    else if (a === '--origin') opts.origin = (argv[++i] ?? '').toUpperCase();
    else if (a === '--dest') opts.dest = (argv[++i] ?? '').toUpperCase();
    else if (a === '--depart') opts.depart = argv[++i] ?? '';
    else if (a === '--return') opts.ret = argv[++i] ?? '';
    else if (a === '--page') opts.page = Number(argv[++i]);
    else if (a === '--page-size') opts.pageSize = Number(argv[++i]);
    else if (a === '--file') opts.file = argv[++i];
  }
  return opts;
}

function toResult(raw, opts) {
  const flights = [];
  for (const f of raw.AvailableFlights ?? []) {
    let price = '';
    if (f.FlightFares?.length) {
      price = String(f.FlightFares[0].Prices?.Adult?.Price ?? '');
    }
    const fn = `${f.CarrierCode ?? 'U2'} ${f.FlightNumber ?? ''}`.trim();
    flights.push({
      id: `${f.DepartureIata}-${f.ArrivalIata}-${opts.depart}`,
      airline: 'easyJet',
      flight_number: fn,
      origin: f.DepartureIata,
      destination: f.ArrivalIata,
      depart_at: f.LocalDepartureTime ?? '',
      arrive_at: f.LocalArrivalTime ?? '',
      stops: 0,
      price,
      currency: 'EUR',
      booking_url: 'https://www.easyjet.com/es/',
    });
  }
  const total = flights.length;
  const start = (opts.page - 1) * opts.pageSize;
  return {
    query: `${opts.origin}-${opts.dest} ${opts.depart}`,
    origin: opts.origin,
    destination: opts.dest,
    depart_date: opts.depart,
    return_date: opts.ret,
    total,
    page: opts.page,
    page_size: opts.pageSize,
    has_next_page: total > opts.page * opts.pageSize,
    flights: flights.slice(start, start + opts.pageSize),
    brand: 'easyJet',
    source: 'ejavailability-browser-mcp',
  };
}

const opts = parseArgs(process.argv);
const body = opts.stdin ? readFileSync(0, 'utf8') : readFileSync(opts.file, 'utf8');
process.stdout.write(JSON.stringify(toResult(JSON.parse(body), opts), null, 2) + '\n');
