package serpapi

// New will return a new instance of SerpAPI
func New(apiKey string) *SerpAPI {
	var s SerpAPI
	s.apiKey = apiKey
	return &s
}

// SerpAPI manages the connection to the SerpAPI service
type SerpAPI struct {
	apiKey string
}
