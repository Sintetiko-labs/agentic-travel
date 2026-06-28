package client

// Availability returns price snapshot from hotel detail (live rooms need booking session).
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {
	view, err := c.Read(hotelID)
	if err != nil {
		return nil, err
	}
	return &AvailSummary{
		CheckIn:  checkIn,
		CheckOut: checkOut,
		Guests:   guests,
		Rooms:    rooms,
		Status:   "from_detail",
		From:     view.Price.Price,
		Currency: view.Price.Currency,
	}, nil
}
