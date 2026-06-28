package client

import "fmt"

// Availability checks room availability (stub).
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {
	return nil, fmt.Errorf("availability not yet implemented for Mama Shelter (hotel=%q)", hotelID)
}
