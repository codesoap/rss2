package rss2

import (
	"encoding/xml"
	"fmt"
)

// Enclosure represents an Item's enclosure element. All attributes must
// be present. Type must be a MIME type.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  int      `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

// NewEnclosure creates a new Enclosure element.
func NewEnclosure(url string, length int, t string) (*Enclosure, error) {
	if len(url) == 0 || len(t) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewEnclosure()`)
	}
	return &Enclosure{
		XMLName: xml.Name{Local: `enclosure`},
		URL:     url,
		Length:  length,
		Type:    t,
	}, nil
}
