package rss2

import (
	"encoding/xml"
	"fmt"
)

// Source represents an Item's source element. URL must be present.
type Source struct {
	XMLName xml.Name `xml:"source"`
	Value   string   `xml:",chardata"`
	URL     string   `xml:"url,attr"`
}

// NewSource creates a new Source element.
func NewSource(value, url string) (*Source, error) {
	if len(value) == 0 || len(url) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewSource()`)
	}
	return &Source{
		XMLName: xml.Name{Local: `source`},
		Value:   value,
		URL:     url,
	}, nil
}
