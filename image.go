package rss2

import (
	"encoding/xml"
	"fmt"
)

// An rss channel's image. URL, Title and Link must be present.
// Width must not exceed 144. Height must not exceed 400.
type Image struct {
	XMLName     xml.Name `xml:"image"`
	URL         string   `xml:"url"`
	Title       string   `xml:"title"`
	Link        string   `xml:"link"`
	Width       int      `xml:"width,omitempty"`
	Height      int      `xml:"height,omitempty"`
	Description string   `xml:"description,omitempty"`
}

// Create new rss image.
func NewImage(url, title, link string) (*Image, error) {
	if len(url) == 0 || len(title) == 0 || len(link) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewImage().`)
	}
	return &Image{
		XMLName: xml.Name{Local: `image`},
		URL:     url,
		Title:   title,
		Link:    link,
	}, nil
}
