package serpapi

// Response is the SerpAPI response
type Response struct {
	SearchMetadata    SearchMetaData    `json:"search_metadata"`
	SearchParameters  SearchParameters  `json:"search_parameters"`
	SearchInformation SearchInformation `json:"search_information"`
	LocalMap          *LocalMap         `json:"local_map"`
	LocalResults      []LocalResult     `json:"local_results"`
	KnowledgeGraph    *KnowledgeGraph   `json:"knowledge_graph"`
	RelatedQuestions  []RelatedQuestion `json:"related_questions"`
	OrganicResults    []OrganicResult   `json:"organic_results"`
	RelatedSearches   []RelatedSearch   `json:"related_searches"`
	Pagination        Pagination        `json:"pagination"`
}
