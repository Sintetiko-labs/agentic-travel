package client

import (
	"fmt"
	"strings"
)

// Read returns flight detail by id or booking URL (uses last search shape).
func (c *Client) Read(idOrURL string) (*FlightView, error) {
	id := strings.TrimSpace(idOrURL)
	if id == "" {
		return nil, fmt.Errorf("id or url required")
	}
	parts := strings.Split(strings.ReplaceAll(id, " ", "-"), "-")
	if len(parts) < 3 {
		return nil, fmt.Errorf("read expects id ORIGIN-DEST-DATE or a booking URL")
	}
	origin, dest, depart := parts[0], parts[1], parts[2]
	res, err := c.Search(origin, dest, depart, "", 1, 1)
	if err != nil {
		return nil, err
	}
	if len(res.Flights) == 0 {
		return nil, fmt.Errorf("no flight found for %q", id)
	}
	f := res.Flights[0]
	return &FlightView{
		ID:           f.ID,
		Airline:      f.Airline,
		FlightNumber: f.FlightNumber,
		Origin:       f.Origin,
		Destination:  f.Destination,
		Depart:       f.Depart,
		Arrive:       f.Arrive,
		Duration:     f.Duration,
		Stops:        f.Stops,
		Price:        PriceInfo{Price: f.Price, Currency: f.Currency},
		Cabin:        f.Cabin,
		BookingURL:   f.BookingURL,
	}, nil
}
