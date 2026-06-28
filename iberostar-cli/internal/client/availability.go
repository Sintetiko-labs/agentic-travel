package client

import "fmt"

// Availability queries Iberostar booking availability.
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {
	payload := map[string]any{
		"query": `query Avail($id:ID!,$in:String!,$out:String!,$g:Int!,$r:Int!){hotelAvailability(hotelId:$id,checkIn:$in,checkOut:$out,guests:$g,rooms:$r){status fromPrice currency}}`,
		"variables": map[string]any{
			"id": urlPathID(hotelID), "in": checkIn, "out": checkOut, "g": guests, "r": rooms,
		},
	}
	var resp struct {
		Data struct {
			HotelAvailability struct {
				Status, Currency string
				FromPrice        float64 `json:"fromPrice"`
			} `json:"hotelAvailability"`
		} `json:"data"`
	}
	if err := c.PostJSON(c.BaseURL+"/api/graphql", payload, &resp); err != nil {
		return nil, fmt.Errorf("availability: %w — try IBEROSTAR_COOKIE", err)
	}
	a := resp.Data.HotelAvailability
	from := ""
	if a.FromPrice > 0 {
		from = fmt.Sprintf("%.2f", a.FromPrice)
	}
	return &AvailSummary{
		CheckIn: checkIn, CheckOut: checkOut, Guests: guests, Rooms: rooms,
		Status: a.Status, From: from, Currency: a.Currency,
	}, nil
}
