package rss2

import (
	"encoding/xml"
	"fmt"
)

// An rss channel's or item's category. Domain is an optional attribute.
type Category struct {
	XMLName xml.Name `xml:"category"`
	Value   string   `xml:",chardata"`
	Domain  string   `xml:"domain,attr,omitempty"`
}

// Create new category rss element.
func NewCategory(value string) (*Category, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewCategory()`)
	}
	return &Category{
		XMLName: xml.Name{Local: `category`},
		Value:   value,
	}, nil
}
