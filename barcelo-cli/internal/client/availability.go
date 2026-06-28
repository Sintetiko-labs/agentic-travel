package client

import (
	"fmt"
	"time"

	tkbase "github.com/fbelchi/travelkit/base"
)

// Availability returns a stub snapshot (Barceló requires booking engine session for live rates).
func (c *Client) Availability(hotelID, checkIn, checkOut string, guests, rooms int) (*AvailSummary, error) {
	if checkIn == "" || checkOut == "" {
		return nil, fmt.Errorf("check-in and check-out dates required")
	}
	view, err := c.Read(hotelID)
	if err != nil {
		return nil, err
	}
	return &AvailSummary{
		CheckIn:  checkIn,
		CheckOut: checkOut,
		Guests:   guests,
		Rooms:    rooms,
		Status:   "check_booking_engine",
		From:     view.Price.Price,
		Currency: view.Price.Currency,
	}, nil
}

func bookingURL(hotelURL, checkIn, checkOut string, guests, rooms int) string {
	if hotelURL == "" {
		return ""
	}
	if guests < 1 {
		guests = 2
	}
	if rooms < 1 {
		rooms = 1
	}
	_ = time.Now()
	return tkbase.Absolutize(hotelURL, fmt.Sprintf("?checkin=%s&checkout=%s&rooms=%d&adults=%d", checkIn, checkOut, rooms, guests))
}
