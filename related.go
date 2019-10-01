package serpapi

// RelatedQuestion represents a related question entry
type RelatedQuestion struct {
	Question      string `json:"question"`
	Snippet       string `json:"snippet"`
	Title         string `json:"title"`
	Link          string `json:"link"`
	DisplayedLink string `json:"displayed_link"`
}

// RelatedSearch represents a related search entry
type RelatedSearch struct {
	Query string `json:"query"`
	Link  string `json:"link"`
}
