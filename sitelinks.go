package serpapi

// Sitelinks represents the sitelinks
type Sitelinks struct {
	Inline []SitelinksEntry `json:"inline"`
}

// SitelinksEntry represents a sitelinks entry
type SitelinksEntry struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}
