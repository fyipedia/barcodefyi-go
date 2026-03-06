// Package barcodefyi provides a Go client for the BarcodeFYI API.
//
// BarcodeFYI is a comprehensive barcode reference covering symbologies, standards,
// components, and industry applications. This client requires no authentication
// and has zero external dependencies.
//
// Usage:
//
//	client := barcodefyi.NewClient()
//	result, err := client.Search("upc")
package barcodefyi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// DefaultBaseURL is the default base URL for the BarcodeFYI API.
const DefaultBaseURL = "https://barcodefyi.com/api"

// Client is a BarcodeFYI API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new BarcodeFYI API client with default settings.
func NewClient() *Client {
	return &Client{
		BaseURL:    DefaultBaseURL,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) get(path string, result interface{}) error {
	resp, err := c.HTTPClient.Get(c.BaseURL + path)
	if err != nil {
		return fmt.Errorf("barcodefyi: request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("barcodefyi: HTTP %d: %s", resp.StatusCode, string(body))
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("barcodefyi: decode failed: %w", err)
	}
	return nil
}

// Search searches across barcode symbologies, standards, and glossary terms.
func (c *Client) Search(query string) (*SearchResult, error) {
	var result SearchResult
	path := fmt.Sprintf("/search/?q=%s", url.QueryEscape(query))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Symbology returns details for a barcode symbology by slug.
func (c *Client) Symbology(slug string) (*SymbologyDetail, error) {
	var result SymbologyDetail
	if err := c.get("/symbology/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Family returns details for a barcode family by slug.
func (c *Client) Family(slug string) (*FamilyDetail, error) {
	var result FamilyDetail
	if err := c.get("/family/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Standard returns details for a barcode standard by slug.
func (c *Client) Standard(slug string) (*StandardDetail, error) {
	var result StandardDetail
	if err := c.get("/standard/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Component returns details for a barcode component by slug.
func (c *Client) Component(slug string) (*ComponentDetail, error) {
	var result ComponentDetail
	if err := c.get("/component/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GlossaryTerm returns a glossary term by slug.
func (c *Client) GlossaryTerm(slug string) (*GlossaryTerm, error) {
	var result GlossaryTerm
	if err := c.get("/glossary/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Compare compares two barcode symbologies.
func (c *Client) Compare(slugA, slugB string) (*CompareResult, error) {
	var result CompareResult
	path := fmt.Sprintf("/compare/?a=%s&b=%s", url.QueryEscape(slugA), url.QueryEscape(slugB))
	if err := c.get(path, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Random returns a random barcode symbology.
func (c *Client) Random() (*SymbologyDetail, error) {
	var result SymbologyDetail
	if err := c.get("/random/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Industry returns details for a barcode industry application by slug.
func (c *Client) Industry(slug string) (*IndustryDetail, error) {
	var result IndustryDetail
	if err := c.get("/industry/"+slug+"/", &result); err != nil {
		return nil, err
	}
	return &result, nil
}
