package rss2

import (
	"testing"
	"time"
)

func TestParseRSSTime(t *testing.T) {
	testCases := map[string]RSSTime{
		"Sat, 07 Sep 2002 00:08:01 UT":  RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Mon, 09 Sep 2002 00:08:01 UT":  RSSTime{time.Date(2002, 9, 9, 0, 8, 1, 0, time.UTC)},
		"Sat, 07 Sep 2002 00:08 UT":     RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"Sat, 07 Sep 2002 00:08:01 GMT": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Sat, 07 Sep 02 00:08:01 GMT":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Sat, 07 Sep 02 00:08 GMT":      RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 02 00:08 GMT":           RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 97 00:08 GMT":           RSSTime{time.Date(1997, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 2002 00:08:01 UT":       RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 00:08 UT":          RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 2002 00:08:01 GMT":      RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Aug 2001 05:08:01 GMT":      RSSTime{time.Date(2001, 8, 7, 5, 8, 1, 0, time.UTC)},
		// Military time zones:
		"07 Sep 2002 00:08:01 Z":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 01:08:01 A":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 02:08:01 B":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 12:08:01 M":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 23:08:01 N":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 12:08:01 Y":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Fri, 06 Sep 2002 12:08 Y": RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		// 3-character timezone indicators for North America:
		"06 Sep 2002 19:08:01 EST": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 20:08:01 EDT": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 18:08:01 CST": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 19:08:01 CDT": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 17:08:01 MST": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 18:08:01 MDT": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 16:08:01 PST": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 17:08:01 PDT": RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 17:08 PDT":    RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		// Explicit indication of the offset from UTC:
		"06 Sep 2002 17:08:01 -0700":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 11:08:01 +1100":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 00:08:01 -0000":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 00:13:01 +0005":   RSSTime{time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Fri, 06 Sep 2002 17:08 -0700": RSSTime{time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
	}
	for in, expected := range testCases {
		if out, err := parseRSSTime(in); err != nil {
			t.Errorf("Error parsing '%s': %s", in, err.Error())
		} else if !out.Time.Equal(expected.Time) {
			t.Errorf("Parsing '%s' yielded '%s'. Expected '%s'", in, out.Time.String(),
				expected.Time.String())
		}
	}
}

func BenchmarkParseRSSTimeEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseRSSTime("06 Sep 2002 17:08:00 +0000")
	}
}

func BenchmarkParseRSSTimeHard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseRSSTime("Mon, 06 Sep 02 17:08 GMT")
	}
}
