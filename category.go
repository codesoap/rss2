package rss2

import (
	"encoding/xml"
	"fmt"
)

// Category represents a Channel's or Item's category. Domain is
// optional.
type Category struct {
	XMLName xml.Name `xml:"category"`
	Value   string   `xml:",chardata"`
	Domain  string   `xml:"domain,attr,omitempty"`
}

// NewCategory creates a new Category.
func NewCategory(value string) (*Category, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewCategory()`)
	}
	return &Category{
		XMLName: xml.Name{Local: `category`},
		Value:   value,
	}, nil
}
