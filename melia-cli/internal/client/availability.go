package client

import "fmt"

// Availability queries Meliá availability BFF (requires session cookie).
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {
	payload := map[string]any{
		"hotelCode": hotelID,
		"checkIn":   checkIn,
		"checkOut":  checkOut,
		"guests":    guests,
		"rooms":     rooms,
		"language":  "es",
	}
	var resp struct {
		Status   string  `json:"status"`
		From     float64 `json:"fromPrice"`
		Currency string  `json:"currency"`
	}
	if err := c.PostJSON(c.BaseURL+"/services/booking/hotels/v1/availability", payload, &resp); err != nil {
		return nil, fmt.Errorf("availability %q: %w — try MELIA_COOKIE", hotelID, err)
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
