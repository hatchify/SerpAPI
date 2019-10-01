package serpapi

// Pagination represents the search results pagination
type Pagination struct {
	Current    int        `json:"current"`
	Next       string     `json:"next"`
	OtherPages OtherPages `json:"other_pages"`
}

// OtherPages represents all the pages available within the pagination
type OtherPages struct {
	Num2  string `json:"2"`
	Num3  string `json:"3"`
	Num4  string `json:"4"`
	Num5  string `json:"5"`
	Num6  string `json:"6"`
	Num7  string `json:"7"`
	Num8  string `json:"8"`
	Num9  string `json:"9"`
	Num10 string `json:"10"`
}
