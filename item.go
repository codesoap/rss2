package rss2

import (
	"encoding/xml"
	"fmt"
)

// Item represents an rss item. At least Title or Description must be
// present.
type Item struct {
	XMLName     xml.Name    `xml:"item"`
	Title       string      `xml:"title,omitempty"`
	Link        string      `xml:"link,omitempty"`
	Description string      `xml:"description,omitempty"`
	Author      string      `xml:"author,omitempty"`
	Categories  []*Category `xml:"category,omitempty"`
	Comments    string      `xml:"comments,omitempty"`
	Enclosure   *Enclosure  `xml:"enclosure,omitempty"`
	GUID        *GUID       `xml:"guid,omitempty"`
	PubDate     *RSSTime    `xml:"pubDate,omitempty"`
	Source      *Source     `xml:"source,omitempty"`
}

// NewItem creates a new Item. Either title or description may be empty.
func NewItem(title, description string) (*Item, error) {
	if len(title) == 0 && len(description) == 0 {
		return nil, fmt.Errorf(`cannot create item with empty title and description`)
	}
	return &Item{
		XMLName:     xml.Name{Local: `item`},
		Title:       title,
		Description: description,
	}, nil
}
