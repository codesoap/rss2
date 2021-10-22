package rss2

import (
	"testing"
	"time"
)

func TestParseRSSTime(t *testing.T) {
	testCases := map[string]RSSTime{
		"Sat, 07 Sep 2002 00:08:01 UT":  {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Mon, 09 Sep 2002 00:08:01 UT":  {time.Date(2002, 9, 9, 0, 8, 1, 0, time.UTC)},
		"Sat, 07 Sep 2002 00:08 UT":     {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"Sat, 07 Sep 2002 00:08:01 GMT": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Sat, 07 Sep 02 00:08:01 GMT":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Sat, 07 Sep 02 00:08 GMT":      {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 02 00:08 GMT":           {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 97 00:08 GMT":           {time.Date(1997, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 2002 00:08:01 UT":       {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 00:08 UT":          {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		"07 Sep 2002 00:08:01 GMT":      {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Aug 2001 05:08:01 GMT":      {time.Date(2001, 8, 7, 5, 8, 1, 0, time.UTC)},
		// Military time zones:
		"07 Sep 2002 00:08:01 Z":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 01:08:01 A":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 02:08:01 B":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 12:08:01 M":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 23:08:01 N":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 12:08:01 Y":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Fri, 06 Sep 2002 12:08 Y": {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		// 3-character timezone indicators for North America:
		"06 Sep 2002 19:08:01 EST": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 20:08:01 EDT": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 18:08:01 CST": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 19:08:01 CDT": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 17:08:01 MST": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 18:08:01 MDT": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 16:08:01 PST": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 17:08:01 PDT": {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"06 Sep 2002 17:08 PDT":    {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
		// Explicit indication of the offset from UTC:
		"06 Sep 2002 17:08:01 -0700":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 11:08:01 +1100":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 00:08:01 -0000":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"07 Sep 2002 00:13:01 +0005":   {time.Date(2002, 9, 7, 0, 8, 1, 0, time.UTC)},
		"Fri, 06 Sep 2002 17:08 -0700": {time.Date(2002, 9, 7, 0, 8, 0, 0, time.UTC)},
	}
	for in, expected := range testCases {
		if out, err := ParseRSSTime(in); err != nil {
			t.Errorf("Error parsing '%s': %s", in, err.Error())
		} else if !out.Time.Equal(expected.Time) {
			t.Errorf("Parsing '%s' yielded '%s'. Expected '%s'", in, out.Time.String(),
				expected.Time.String())
		}
	}
}

func BenchmarkParseRSSTimeEasy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseRSSTime("06 Sep 2002 17:08:00 +0000")
	}
}

func BenchmarkParseRSSTimeHard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseRSSTime("Mon, 06 Sep 02 17:08 GMT")
	}
}
