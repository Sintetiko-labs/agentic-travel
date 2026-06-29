#!/usr/bin/env node
/**
 * Merge parallel wave source JSON files into CombinedSearchResult.
 * Reads manifest from stdin: { query, sources: [{id, status, duration_ms, file?, error?}], mcp_agent_fallback? }
 */
import fs from 'node:fs';

function readJson(path) {
  if (!path || !fs.existsSync(path)) return null;
  try {
    return JSON.parse(fs.readFileSync(path, 'utf8'));
  } catch {
    return null;
  }
}

function asArray(v) {
  return Array.isArray(v) ? v : [];
}

function pushFlights(out, data, source) {
  if (!data) return 0;
  if (Array.isArray(data.flights)) {
    for (const f of data.flights) out.push({ ...f, source: f.source ?? source });
    return data.flights.length;
  }
  if (Array.isArray(data.offers)) {
    for (const o of data.offers) {
      out.push({
        id: o.id ?? o.offer_id ?? '',
        airline: o.carrier ?? o.airline ?? '',
        origin: o.origin ?? o.from ?? '',
        destination: o.destination ?? o.to ?? '',
        depart_at: o.departure ?? o.depart_at ?? '',
        arrive_at: o.arrival ?? o.arrive_at ?? '',
        stops: o.stops ?? 0,
        price: String(o.price ?? o.total_amount ?? ''),
        currency: o.currency ?? '',
        booking_url: o.booking_url ?? o.deep_link ?? '',
        source,
      });
    }
    return data.offers.length;
  }
  return 0;
}

function pushHotels(out, data, source) {
  if (!data) return 0;
  if (Array.isArray(data.hotels)) {
    for (const h of data.hotels) out.push({ ...h, source: h.source ?? source });
    return data.hotels.length;
  }
  if (Array.isArray(data.results)) {
    for (const h of data.results) {
      out.push({
        id: h.id ?? h.hotel_id ?? '',
        name: h.name ?? h.hotel_name ?? '',
        city: h.city ?? '',
        price: String(h.price ?? h.rate ?? ''),
        currency: h.currency ?? '',
        hotel_url: h.url ?? h.booking_url ?? '',
        source,
      });
    }
    return data.results.length;
  }
  return 0;
}

const manifest = JSON.parse(fs.readFileSync(0, 'utf8'));
const flights = [];
const hotels = [];
const sources = [];
const timedOut = [];

for (const src of manifest.sources ?? []) {
  const data = src.file ? readJson(src.file) : null;
  let flightCount = 0;
  let hotelCount = 0;
  if (data && src.status === 'ok') {
    flightCount = pushFlights(flights, data, src.id);
    hotelCount = pushHotels(hotels, data, src.id);
  }
  sources.push({
    id: src.id,
    status: src.status,
    duration_ms: src.duration_ms ?? 0,
    ...(flightCount ? { flights: flightCount } : {}),
    ...(hotelCount ? { hotels: hotelCount } : {}),
    ...(src.error ? { error: src.error } : {}),
  });
  if (src.status === 'timed_out') timedOut.push(src.id);
}

const result = {
  query: manifest.query ?? {},
  flights,
  hotels,
  sources,
  timed_out: timedOut,
  wall_clock_ms: manifest.wall_clock_ms ?? 0,
};
if (manifest.mcp_agent_fallback?.length) {
  result.mcp_agent_fallback = manifest.mcp_agent_fallback;
}

console.log(JSON.stringify(result, null, 2));
