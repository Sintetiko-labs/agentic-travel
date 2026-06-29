package types

// WaveSourceStatus is per-source timing from a parallel wave search.
type WaveSourceStatus struct {
	ID         string `json:"id"`
	Status     string `json:"status"` // ok | error | timed_out | skipped
	DurationMS int64  `json:"duration_ms"`
	Flights    int    `json:"flights,omitempty"`
	Hotels     int    `json:"hotels,omitempty"`
	Error      string `json:"error,omitempty"`
}

// WaveQuery captures the wave search parameters.
type WaveQuery struct {
	From   string `json:"from,omitempty"`
	To     string `json:"to,omitempty"`
	City   string `json:"city,omitempty"`
	Depart string `json:"depart,omitempty"`
	CheckIn  string `json:"check_in,omitempty"`
	CheckOut string `json:"check_out,omitempty"`
}

// MCPAgentFallback documents MCP calls the agent should run when HTTP client fails.
type MCPAgentFallback struct {
	Server string         `json:"server"`
	Tool   string         `json:"tool"`
	Args   map[string]any `json:"args,omitempty"`
	Note   string         `json:"note,omitempty"`
}

// CombinedSearchResult is the merged output of MCP + CLI parallel wave search.
// flights[] and hotels[] are never null.
type CombinedSearchResult struct {
	Query            WaveQuery            `json:"query"`
	Flights          []FlightHit          `json:"flights"`
	Hotels           []HotelHit           `json:"hotels"`
	Sources          []WaveSourceStatus   `json:"sources"`
	TimedOut         []string             `json:"timed_out"`
	MCPAgentFallback []MCPAgentFallback   `json:"mcp_agent_fallback,omitempty"`
	WallClockMS      int64                `json:"wall_clock_ms,omitempty"`
}
