package rss2_test

import (
	"encoding/xml"
	"fmt"

	"github.com/codesoap/rss2"
)

// This example shows how you can parse a simple RSS 2.0 feed.
// The user must ensure that the input is UTF-8 encoded when parsing.
func Example_parseRSS() {
	input := `
		<?xml version="1.0" encoding="UTF-8"?>
		<rss version="2.0">
		    <channel>
		        <title>Channel title</title>
		        <link>channel.link.net</link>
		        <description>Channel description</description>
		        <category>Channel&#39;s category</category>
		        <item>
		            <title>Title of my RSS item</title>
		            <pubDate>03 Jun 2019 09:39:21 GMT</pubDate>
		        </item>
		    </channel>
		</rss>`
	var parse rss2.RSS
	// Remember to handle any error xml.Unmarshal() gives in your code.
	xml.Unmarshal([]byte(input), &parse)
	fmt.Println(`RSS Version:        `, parse.Version)
	fmt.Println(`RSS channel's title:`, parse.Channel.Title)
	fmt.Println(`RSS item's title:   `, parse.Channel.Items[0].Title)
	// The date of any date fields is hidden behind .Time for technical reasons.
	fmt.Println(`RSS item's pubDate: `, parse.Channel.Items[0].PubDate.Time)
}
