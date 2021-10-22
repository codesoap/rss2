package rss2

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var reDayName = regexp.MustCompile(`^(Mon|Tue|Wed|Thu|Fri|Sat|Sun), `)
var reSeconds = regexp.MustCompile(` ([0-2][0-9]:[0-5][0-9]) `)
var reYear = regexp.MustCompile(`^([0-3][0-9] [A-Za-z]{3} )([0-9]{2})( .*)$`)
var reTimezone = regexp.MustCompile(`^(.* )([A-Z]{1,3})$`)

// RSSTime is a wrapper around time.Time, that makes it possible to
// define custom MarshalXML() and UnmarshalXML() functions.
type RSSTime struct {
	Time time.Time
}

// UnmarshalXML unmarshals an RSSTime element.
func (t *RSSTime) UnmarshalXML(decoder *xml.Decoder,
	start xml.StartElement) (err error) {
	var value string
	if err = decoder.DecodeElement(&value, &start); err == nil {
		*t, err = ParseRSSTime(value)
	}
	return
}

// MarshalXML marshals an RSSTime element.
func (t RSSTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(t.Time.Format("02 Jan 2006 15:04:05 -0700"), start)
}

// ParseRSSTime parses a time as specified in RFC822, with the
// difference, that four digit years are allowed.
func ParseRSSTime(in string) (out RSSTime, err error) {
	tmp := []byte(in)
	// TODO: Research whether multiple consecutive whitspaces are legal
	//       and reduce them to single ' ', if so.
	tmp = removeDayNameIfPresent(tmp)
	tmp = addSecondsIfMissing(tmp)
	if tmp, err = convertToFourDigitYearIfNeeded(tmp); err != nil {
		return
	}
	if tmp, err = convertToNumericTimezoneIfNeeded(tmp); err != nil {
		return
	}
	t, err := time.Parse("02 Jan 2006 15:04:05 -0700", string(tmp))
	return RSSTime{t}, err
}

func removeDayNameIfPresent(in []byte) []byte {
	return reDayName.ReplaceAllLiteral(in, nil)
}

func addSecondsIfMissing(in []byte) []byte {
	return reSeconds.ReplaceAll(in, []byte(` ${1}:00 `))
}

func convertToFourDigitYearIfNeeded(in []byte) (r []byte, err error) {
	twoDigitYear := reYear.ReplaceAll(in, []byte(`$2`))
	if len(twoDigitYear) == 2 {
		var fourDigitYear []byte
		var yearSuffix int
		yearSuffix, err = strconv.Atoi(string(twoDigitYear))
		if err != nil {
			return
		}
		if yearSuffix >= 70 {
			fourDigitYear = append([]byte("19"), twoDigitYear...)
		} else {
			fourDigitYear = append([]byte("20"), twoDigitYear...)
		}
		repl := append(append([]byte(`${1}`), fourDigitYear...), `${3}`...)
		r = reYear.ReplaceAll(in, repl)
	} else {
		r = in
	}
	return
}

func convertToNumericTimezoneIfNeeded(in []byte) (r []byte, err error) {
	// The values are the offset from UTC in hours:
	timezones := map[string]int{
		"UT": 0, "GMT": 0,
		"EST": -5, "EDT": -4,
		"CST": -6, "CDT": -5,
		"MST": -7, "MDT": -6,
		"PST": -8, "PDT": -7,
		"A": 1, "B": 2, "C": 3, "D": 4, "E": 5, "F": 6, "G": 7, "H": 8, "I": 9,
		"K": 10, "L": 11, "M": 12, "N": -1, "O": -2, "P": -3, "Q": -4, "R": -5,
		"S": -6, "T": -7, "U": -8, "V": -9, "W": -10, "X": -11, "Y": -12, "Z": 0,
	}

	timezone := reTimezone.ReplaceAll(in, []byte(`$2`))
	if len(timezone) > 0 && len(timezone) <= 3 {
		utcOffsetInHours, ok := timezones[string(timezone)]
		if !ok {
			return nil, fmt.Errorf("invalid timezone '%s'", string(timezone))
		}
		numericTimezone := fmt.Sprintf("%+03d00", utcOffsetInHours)
		r = reTimezone.ReplaceAll(in, append([]byte(`${1}`), numericTimezone...))
	} else {
		r = in
	}
	return
}
