package client

import "fmt"

func paginateFlightResult(origin, dest, depart, ret string, page, pageSize int, flights []FlightHit, brand, source string) *FlightSearchResult {
	total := len(flights)
	start := (page - 1) * pageSize
	if start > total { start = total }
	end := start + pageSize
	if end > total { end = total }
	pageFlights := flights[start:end]
	if pageFlights == nil { pageFlights = []FlightHit{} }
	return &FlightSearchResult{Query: fmt.Sprintf("%s-%s %s", origin, dest, depart), Origin: origin, Dest: dest, Depart: depart, Return: ret, Total: total, Page: page, PageSize: pageSize, HasNext: total > page*pageSize, Flights: pageFlights, Brand: brand, Source: source}
}
