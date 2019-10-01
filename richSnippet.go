package serpapi

// RichSnippet represent an organic result rich snippet
type RichSnippet struct {
	Bottom RichSnippetBottom `json:"bottom"`
}

// RichSnippetBottom represent an organic result rich snippet bottom
type RichSnippetBottom struct {
	Extensions []string `json:"extensions"`
}
