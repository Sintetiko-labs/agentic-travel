package types

// HotelSearchResult is normalized hotel search output.
type HotelSearchResult struct {
	Query    string      `json:"query"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	HasNext  bool        `json:"has_next_page"`
	Hotels   []HotelHit  `json:"hotels"`
	Brand    string      `json:"brand,omitempty"`
	Source   string      `json:"source"`
}

// HotelHit is a compact hotel row from search.
type HotelHit struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand,omitempty"`
	City        string  `json:"city,omitempty"`
	Country     string  `json:"country,omitempty"`
	Stars       float64 `json:"stars,omitempty"`
	Price       string  `json:"price,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	Rating      float64 `json:"rating,omitempty"`
	ReviewCount int     `json:"review_count,omitempty"`
	HotelURL    string  `json:"hotel_url"`
	ImageURL    string  `json:"image_url,omitempty"`
}

// HotelView is projected hotel detail.
type HotelView struct {
	ID           string         `json:"id"`
	Name         string         `json:"name"`
	Brand        string         `json:"brand,omitempty"`
	Description  string         `json:"description,omitempty"`
	Address      string         `json:"address,omitempty"`
	City         string         `json:"city,omitempty"`
	Country      string         `json:"country,omitempty"`
	Stars        float64        `json:"stars,omitempty"`
	HotelURL     string         `json:"hotel_url"`
	Price        PriceInfo      `json:"price"`
	Amenities    []string       `json:"amenities,omitempty"`
	Availability *AvailSummary  `json:"availability,omitempty"`
}

// AvailSummary is a compact availability snapshot.
type AvailSummary struct {
	CheckIn  string `json:"check_in,omitempty"`
	CheckOut string `json:"check_out,omitempty"`
	Rooms    int    `json:"rooms,omitempty"`
	Guests   int    `json:"guests,omitempty"`
	Status   string `json:"status,omitempty"`
	From     string `json:"from_price,omitempty"`
	Currency string `json:"currency,omitempty"`
}

// PriceInfo is shared price structure.
type PriceInfo struct {
	Price         string `json:"price"`
	OriginalPrice string `json:"original_price,omitempty"`
	Currency      string `json:"currency,omitempty"`
	PerNight      bool   `json:"per_night,omitempty"`
}

// FlightSearchResult is normalized flight search output.
type FlightSearchResult struct {
	Query    string      `json:"query"`
	Origin   string      `json:"origin,omitempty"`
	Dest     string      `json:"destination,omitempty"`
	Depart   string      `json:"depart_date,omitempty"`
	Return   string      `json:"return_date,omitempty"`
	Total    int         `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	HasNext  bool        `json:"has_next_page"`
	Flights  []FlightHit `json:"flights"`
	Brand    string      `json:"brand,omitempty"`
	Source   string      `json:"source"`
}

// FlightHit is a compact flight offer row.
type FlightHit struct {
	ID           string `json:"id"`
	Airline      string `json:"airline,omitempty"`
	FlightNumber string `json:"flight_number,omitempty"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Depart       string `json:"depart_at"`
	Arrive       string `json:"arrive_at"`
	Duration     string `json:"duration,omitempty"`
	Stops        int    `json:"stops"`
	Price        string `json:"price"`
	Currency     string `json:"currency,omitempty"`
	Cabin        string `json:"cabin,omitempty"`
	BookingURL   string `json:"booking_url,omitempty"`
}

// FlightView is projected flight or fare detail.
type FlightView struct {
	ID           string    `json:"id"`
	Airline      string    `json:"airline,omitempty"`
	FlightNumber string    `json:"flight_number,omitempty"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	Depart       string    `json:"depart_at"`
	Arrive       string    `json:"arrive_at"`
	Duration     string    `json:"duration,omitempty"`
	Stops        int       `json:"stops"`
	Price        PriceInfo `json:"price"`
	Cabin        string    `json:"cabin,omitempty"`
	BookingURL   string    `json:"booking_url,omitempty"`
	Segments     []Segment `json:"segments,omitempty"`
}

// Segment is one leg of a multi-stop itinerary.
type Segment struct {
	Airline      string `json:"airline,omitempty"`
	FlightNumber string `json:"flight_number,omitempty"`
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Depart       string `json:"depart_at"`
	Arrive       string `json:"arrive_at"`
}
