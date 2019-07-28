package rss2

import (
	"encoding/xml"
	"fmt"
)

// rss channel element. Title, Link and Description are required.
type Channel struct {
	XMLName        xml.Name   `xml:"channel"`
	Title          string     `xml:"title"`
	Link           string     `xml:"link"`
	Description    string     `xml:"description"`
	Language       string     `xml:"language,omitempty"`
	Copyright      string     `xml:"copyright,omitempty"`
	ManagingEditor string     `xml:"managingEditor,omitempty"`
	WebMaster      string     `xml:"webMaster,omitempty"`
	PubDate        *RSSTime   `xml:"pubDate,omitempty"`
	LastBuildDate  *RSSTime   `xml:"lastBuildDate,omitempty"`
	Category       *Category  `xml:"category,omitempty"`
	Generator      string     `xml:"generator,omitempty"`
	Docs           string     `xml:"docs,omitempty"`
	Cloud          *Cloud     `xml:"cloud,omitempty"`
	TTL            int        `xml:"ttl,omitempty"`
	Image          *Image     `xml:"image,omitempty"`
	Rating         string     `xml:"rating,omitempty"`
	TextInput      *TextInput `xml:"textInput,omitempty"`
	SkipHours      *SkipHours `xml:"skipHours,omitempty"`
	SkipDays       *SkipDays  `xml:"skipDays,omitempty"`
	Items          []*Item    `xml:"item,omitempty"`
}

// Create a new rss channel.
func NewChannel(title, link, description string) (*Channel, error) {
	if len(title) == 0 || len(link) == 0 || len(description) == 0 {
		return nil, fmt.Errorf(`empty string passed to NewChannel()`)
	}
	return &Channel{
		XMLName:     xml.Name{Local: `channel`},
		Title:       title,
		Link:        link,
		Description: description,
	}, nil
}
