package barcodefyi

// SearchResult represents the API search response.
type SearchResult struct {
	Query   string       `json:"query"`
	Results []SearchItem `json:"results"`
	Total   int          `json:"total"`
}

// SearchItem represents a single search result item.
type SearchItem struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	Type string `json:"type"`
	URL  string `json:"url"`
}

// SymbologyDetail represents a barcode symbology.
type SymbologyDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Family      string `json:"family"`
	URL         string `json:"url"`
}

// FamilyDetail represents a barcode family.
type FamilyDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// StandardDetail represents a barcode standard.
type StandardDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// ComponentDetail represents a barcode component.
type ComponentDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

// GlossaryTerm represents a glossary term.
type GlossaryTerm struct {
	Term       string `json:"term"`
	Slug       string `json:"slug"`
	Definition string `json:"definition"`
	URL        string `json:"url"`
}

// CompareResult represents a comparison between two symbologies.
type CompareResult struct {
	SymbologyA interface{} `json:"symbology_a"`
	SymbologyB interface{} `json:"symbology_b"`
	URL        string      `json:"url"`
}

// IndustryDetail represents a barcode industry application.
type IndustryDetail struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	URL         string `json:"url"`
}
