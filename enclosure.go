package rss2

import (
	"encoding/xml"
	"fmt"
)

// An rss item's enclosure element. All attributes must be present.
// Type must be a MIME type.
type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  int      `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

// Create new type rss element.
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
