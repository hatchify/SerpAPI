package serpapi

// KnowledgeGraph represents the knowledge graph
type KnowledgeGraph struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Source      struct {
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"source"`
	CountriesOfOrigin   string `json:"countries_of_origin"`
	PeopleAlsoSearchFor []struct {
		Name   string `json:"name"`
		Link   string `json:"link"`
		Source string `json:"source"`
		Image  string `json:"image"`
	} `json:"people_also_search_for"`
	PeopleAlsoSearchForLink string              `json:"people_also_search_for_link"`
	List                    map[string][]string `json:"list"`
}
