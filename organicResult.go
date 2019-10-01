package serpapi

// OrganicResult represents an organic result entry
type OrganicResult struct {
	Position         int          `json:"position"`
	Title            string       `json:"title"`
	Link             string       `json:"link"`
	DisplayedLink    string       `json:"displayed_link"`
	Snippet          string       `json:"snippet"`
	Sitelinks        *Sitelinks   `json:"sitelinks,omitempty"`
	CachedPageLink   string       `json:"cached_page_link,omitempty"`
	RelatedPagesLink string       `json:"related_pages_link,omitempty"`
	Date             string       `json:"date,omitempty"`
	RichSnippet      *RichSnippet `json:"rich_snippet,omitempty"`
}
