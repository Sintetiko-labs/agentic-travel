package types

import "time"

// CombinedQuery describes the fan-out search parameters.
type CombinedQuery struct {
	Kind   string `json:"kind"` // "flights" | "hotels"
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	Depart string `json:"depart,omitempty"`
	Return string `json:"return,omitempty"`
	City   string `json:"city,omitempty"`
}

// SourceResult is per-provider outcome metadata.
type SourceResult struct {
	Slug       string `json:"slug"`
	OK         bool   `json:"ok"`
	DurationMs int64  `json:"duration_ms"`
	Total      int    `json:"total"`
	Error      string `json:"error,omitempty"`
}

// CombinedSearchResult merges parallel CLI JSON into one agent payload.
type CombinedSearchResult struct {
	Query      CombinedQuery  `json:"query"`
	SearchedAt time.Time      `json:"searched_at"`
	DurationMs int64          `json:"duration_ms"`
	Workers    int            `json:"workers"`
	TimeoutSec int            `json:"timeout_sec"`
	Sources    []SourceResult `json:"sources"`
	Flights    []FlightHit    `json:"flights,omitempty"`
	Hotels     []HotelHit     `json:"hotels,omitempty"`
	Total      int            `json:"total"`
}
