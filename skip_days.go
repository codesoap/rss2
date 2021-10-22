package rss2

import (
	"encoding/xml"
	"time"
)

// SkipDays represents a Channel's skipDays element.
type SkipDays struct {
	XMLName xml.Name `xml:"skipDays"`
	Days    []string `xml:"day"`
}

// NewSkipDays creates a new SkipDays element.
func NewSkipDays(days []time.Weekday) *SkipDays {
	dayNames := map[time.Weekday]string{
		time.Sunday:    `Sunday`,
		time.Monday:    `Monday`,
		time.Tuesday:   `Tuesday`,
		time.Wednesday: `Wednesday`,
		time.Thursday:  `Thursday`,
		time.Friday:    `Friday`,
		time.Saturday:  `Saturday`,
	}
	var daysString []string
	for _, day := range days {
		daysString = append(daysString, dayNames[day])
	}
	return &SkipDays{
		XMLName: xml.Name{Local: `skipDays`},
		Days:    daysString,
	}
}
