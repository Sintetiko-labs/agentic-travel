package client

import "fmt"

// Availability queries NH availability API.
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {
	path := fmt.Sprintf("/nh/es/api/v1/hotels/%s/availability?checkIn=%s&checkOut=%s&guests=%d&rooms=%d",
		urlPathID(hotelID), checkIn, checkOut, guests, rooms)
	var resp struct {
		Status   string  `json:"status"`
		From     float64 `json:"fromPrice"`
		Currency string  `json:"currency"`
	}
	if err := c.GetJSON(path, &resp); err != nil {
		return nil, fmt.Errorf("availability: %w — try NH_COOKIE", err)
	}
	from := ""
	if resp.From > 0 {
		from = fmt.Sprintf("%.2f", resp.From)
	}
	return &AvailSummary{
		CheckIn: checkIn, CheckOut: checkOut, Guests: guests, Rooms: rooms,
		Status: resp.Status, From: from, Currency: resp.Currency,
	}, nil
}
