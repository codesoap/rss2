package rss2

import "encoding/xml"

// RSS represents an rss feed. All fields are mandatory.
// Version must be "2.0" for this library.
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *Channel `xml:"channel"`
}

// NewRSS creates a new RSS element.
func NewRSS(ch *Channel) *RSS {
	return &RSS{
		XMLName: xml.Name{Local: `rss`},
		Version: `2.0`,
		Channel: ch,
	}
}
