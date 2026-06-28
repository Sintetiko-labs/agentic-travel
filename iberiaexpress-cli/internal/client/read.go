package client

import (
	"fmt"
	"strings"
)

// Read returns flight detail by id.
func (c *Client) Read(idOrURL string) (*FlightView, error) {
	parts := strings.Split(strings.TrimSpace(idOrURL), "-")
	if len(parts) < 3 {
		return nil, fmt.Errorf("read expects ORIGIN-DEST-DATE id")
	}
	res, err := c.Search(parts[0], parts[1], parts[2], "", 1, 1)
	if err != nil {
		return nil, err
	}
	if len(res.Flights) == 0 {
		return nil, fmt.Errorf("flight not found: %q", idOrURL)
	}
	f := res.Flights[0]
	return &FlightView{
		ID: f.ID, Airline: f.Airline, FlightNumber: f.FlightNumber,
		Origin: f.Origin, Destination: f.Destination,
		Depart: f.Depart, Arrive: f.Arrive,
		Price: PriceInfo{Price: f.Price, Currency: f.Currency},
		BookingURL: f.BookingURL,
	}, nil
}
