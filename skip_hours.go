package rss2

import (
	"encoding/xml"
	"fmt"
)

// rss channel's skipHours element. Hours must be between 0 and 23.
type SkipHours struct {
	XMLName xml.Name `xml:"skipHours"`
	Hours   []int    `xml:"hour"`
}

// Create new skipHours rss element. Hours must be between 0 and 23.
func NewSkipHours(hours []int) (*SkipHours, error) {
	for _, hour := range hours {
		if hour < 0 || hour > 23 {
			return nil, fmt.Errorf(`hour %d not between 0 and 23`, hour)
		}
	}
	return &SkipHours{
		XMLName: xml.Name{Local: `cannel`},
		Hours:   hours,
	}, nil
}
