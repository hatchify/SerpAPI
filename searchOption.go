package serpapi

// NewSearchOption will return a new search option
func NewSearchOption(key, value string) (s SearchOption) {
	s.Key = key
	s.Value = value
	return
}

// SearchOption represent an additional optional search parameter
// Note: See https://serpapi.com/search-api for a list of available options
type SearchOption struct {
	Key   string
	Value string
}
