package serpapi

// SearchInformation represents search metric information
type SearchInformation struct {
	TotalResults       int64   `json:"total_results"`
	TimeTakenDisplayed float64 `json:"time_taken_displayed"`
	QueryDisplayed     string  `json:"query_displayed"`
}

// SearchMetaData represents the search metadata
type SearchMetaData struct {
	ID             string  `json:"id"`
	Status         string  `json:"status"`
	CreatedAt      string  `json:"created_at"`
	ProcessedAt    string  `json:"processed_at"`
	GoogleURL      string  `json:"google_url"`
	TotalTimeTaken float64 `json:"total_time_taken"`
}

// SearchParameters represents the search parameters
type SearchParameters struct {
	Q                 string `json:"q"`
	LocationRequested string `json:"location_requested"`
	LocationUsed      string `json:"location_used"`
	GoogleDomain      string `json:"google_domain"`
	Hl                string `json:"hl"`
	Gl                string `json:"gl"`
	Device            string `json:"device"`
}
