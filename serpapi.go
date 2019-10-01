package serpapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	// Host of the SerpAPI service
	host = "https://serpapi.com"

	// Search endpoint
	searchEndpoint = "search"
)

// New will return a new instance of SerpAPI
func New(apiKey string) (sp *SerpAPI, err error) {
	var s SerpAPI
	if s.host, err = url.Parse(host); err != nil {
		return
	}

	s.apiKey = apiKey
	sp = &s
	return
}

// SerpAPI manages the connection to the SerpAPI service
type SerpAPI struct {
	hc http.Client

	apiKey string
	host   *url.URL
}

func (s *SerpAPI) newURL(endpoint string) (u url.URL) {
	u = *s.host
	u.Path = endpoint
	return
}

func (s *SerpAPI) newQueryParams(query string, opts []SearchOption) (q url.Values) {
	// Pre-allocate query
	q = make(url.Values, len(opts)+1)
	// Set primary query value
	q.Set("q", query)
	// Set API key value
	q.Set("api_key", s.apiKey)

	// Iterate through optional values
	for _, opt := range opts {
		// Set optional key/value
		q.Set(opt.Key, opt.Value)
	}

	return
}

// Search performs a search for a given query and requested options
func (s *SerpAPI) Search(query string, opts ...SearchOption) (rp *Response, err error) {
	u := s.newURL(searchEndpoint)
	q := s.newQueryParams(query, opts)
	u.RawQuery = q.Encode()

	var req *http.Request
	if req, err = http.NewRequest(http.MethodGet, u.String(), nil); err != nil {
		return
	}
	var httpResp *http.Response
	if httpResp, err = s.hc.Do(req); err != nil {
		return
	}
	defer httpResp.Body.Close()

	var resp Response
	if err = json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
		return
	}

	rp = &resp
	return
}
