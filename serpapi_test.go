package serpapi

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	var err error
	if _, err = testInit(); err != nil {
		t.Fatal(err)
	}
}

func TestSerpAPI_Search(t *testing.T) {
	var (
		s   *SerpAPI
		err error
	)

	if s, err = testInit(); err != nil {
		t.Fatal(err)
	}

	var resp *Response
	if resp, err = s.Search("usehatchapp"); err != nil {
		t.Fatal(err)
	}

	if len(resp.OrganicResults) == 0 {
		t.Fatal("no results found when results were expected")
	}

	first := resp.OrganicResults[0]
	if first.Link != "https://usehatchapp.com/" {
		t.Fatalf("invalid link, expected \"%s\" and received \"%s\"", "https://usehatchapp.com/", first.Link)
	}
}

func testInit() (s *SerpAPI, err error) {
	apiKey := os.Getenv("SERPAPI_APIKEY")
	return New(apiKey)
}
