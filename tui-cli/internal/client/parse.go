package client

import ("encoding/json"; "fmt"; "strconv"; "strings")

func flightsFromJSON(body []byte, origin, dest, depart, airline string) []FlightHit {
	var root any
	if err := json.Unmarshal(body, &root); err != nil { return []FlightHit{} }
	found := collectOfferMaps(root)
	hits := make([]FlightHit, 0, len(found)); seen := make(map[string]struct{})
	for _, m := range found {
		h := mapToFlightHit(m, origin, dest, depart, airline)
		if h == nil { continue }
		key := h.ID; if key == "" { key = h.Origin + h.Destination + h.Depart + h.Price }
		if _, ok := seen[key]; ok { continue }
		seen[key] = struct{}{}; hits = append(hits, *h)
	}
	return hits
}
func collectOfferMaps(v any) []map[string]any {
	var out []map[string]any
	walkJSON(v, func(m map[string]any) { if looksLikeOffer(m) { out = append(out, m) } })
	return out
}
func walkJSON(v any, fn func(map[string]any)) {
	switch x := v.(type) {
	case map[string]any: fn(x); for _, ch := range x { walkJSON(ch, fn) }
	case []any: for _, ch := range x { walkJSON(ch, fn) }
	}
}
func looksLikeOffer(m map[string]any) bool {
	price := firstString(m, "price", "amount", "totalAmount", "fare", "lowestPrice", "value", "totalPrice")
	date := firstString(m, "departureDate", "depDate", "date", "outboundDate", "localDepartureTime", "departure")
	orig := firstString(m, "origin", "departure", "from", "departureIata", "departureAirport", "originAirportCode", "departureStation")
	dst := firstString(m, "destination", "arrival", "to", "arrivalIata", "arrivalAirport", "destinationAirportCode", "arrivalStation")
	return (price != "" && (date != "" || orig != "" || dst != "")) || (orig != "" && dst != "" && date != "")
}
func mapToFlightHit(m map[string]any, origin, dest, depart, airline string) *FlightHit {
	o := strings.ToUpper(firstString(m, "origin", "departure", "from", "departureIata", "departureAirport", "originAirportCode", "departureStation"))
	d := strings.ToUpper(firstString(m, "destination", "arrival", "to", "arrivalIata", "arrivalAirport", "destinationAirportCode", "arrivalStation"))
	if o == "" { o = origin }; if d == "" { d = dest }
	date := firstString(m, "departureDate", "depDate", "date", "outboundDate", "localDepartureTime", "departure")
	if date == "" { date = depart }
	price, curr := firstPrice(m)
	fn := firstString(m, "flightNumber", "flightNo", "number")
	if fn == "" { c := firstString(m, "carrierCode", "airlineCode", "marketingCarrier"); n := firstString(m, "flightNum", "flight"); if c != "" && n != "" { fn = c + " " + n } }
	id := fmt.Sprintf("%s-%s-%s", o, d, date); if price != "" { id += "-" + price }
	return &FlightHit{ID: id, Airline: airline, FlightNumber: fn, Origin: o, Destination: d, Depart: date, Arrive: firstString(m, "arrivalTime", "localArrivalTime", "arrival"), Price: price, Currency: curr, BookingURL: firstString(m, "bookingUrl", "deeplink", "url")}
}
func firstString(m map[string]any, keys ...string) string { for _, k := range keys { if v, ok := m[k]; ok { if s := stringify(v); s != "" { return s } } }; return "" }
func firstPrice(m map[string]any) (string, string) {
	for _, k := range []string{"price", "amount", "totalAmount", "fare", "lowestPrice", "value", "totalPrice"} {
		if v, ok := m[k]; ok { switch x := v.(type) {
		case map[string]any: p := firstString(x, "amount", "value", "total"); c := firstString(x, "currency", "currencyCode"); if p != "" { return p, firstNonEmpty(c, "EUR") }
		case float64: return fmt.Sprintf("%.2f", x), "EUR"
		case string: if x != "" { return x, "EUR" }
		} }
	}
	return "", "EUR"
}
func stringify(v any) string { switch x := v.(type) {
case string: return strings.TrimSpace(x)
case float64: if x == float64(int64(x)) { return strconv.FormatInt(int64(x), 10) }; return strconv.FormatFloat(x, 'f', -1, 64)
default: return "" } }
func firstNonEmpty(vals ...string) string { for _, v := range vals { if v != "" { return v } }; return "" }
