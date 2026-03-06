package barcodefyi_test

import (
	"testing"

	barcodefyi "github.com/fyipedia/barcodefyi-go"
)

func TestNewClient(t *testing.T) {
	c := barcodefyi.NewClient()
	if c.BaseURL != barcodefyi.DefaultBaseURL {
		t.Errorf("expected %s, got %s", barcodefyi.DefaultBaseURL, c.BaseURL)
	}
	if c.HTTPClient == nil {
		t.Error("expected non-nil HTTPClient")
	}
}

func TestSearch(t *testing.T) {
	c := barcodefyi.NewClient()
	result, err := c.Search("upc")
	if err != nil {
		t.Fatalf("Search failed: %v", err)
	}
	if result.Total == 0 {
		t.Error("expected results, got 0")
	}
}
