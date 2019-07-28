package rss2

import (
	"encoding/xml"
	"fmt"
)

// An rss channel's guid element. The IsPermaLink attribute is optional.
type GUID struct {
	XMLName     xml.Name `xml:"guid"`
	Value       string   `xml:",chardata"`
	IsPermaLink bool     `xml:"isPermaLink,attr,omitempty"`
}

// Create new guid rss element.
func NewGUID(value string) (*GUID, error) {
	if len(value) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewGUID().`)
	}
	return &GUID{
		XMLName: xml.Name{Local: `guid`},
		Value:   value,
	}, nil
}
