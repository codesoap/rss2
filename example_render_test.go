package rss2_test

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/codesoap/rss2"
)

// This example shows how you can create a simple RSS 2.0 feed.
// Using the New<element>() functions to create RSS elements helps to
// avoid invalid RSS.
func Example_renderRSS() {
	item, _ := rss2.NewItem(`Title of my RSS item`, ``)
	item.PubDate = &rss2.RSSTime{time.Date(2019, 6, 3, 9, 39, 21, 0, time.UTC)}

	category, _ := rss2.NewCategory(`MSFT`)
	category.Domain = `http://www.fool.com/cusips`

	channel, _ := rss2.NewChannel(`Channel title`, `channel.link.net`,
		`Channel description`)
	channel.Categories = []*rss2.Category{category}
	channel.Items = []*rss2.Item{item}

	// All created RSS elements are structs with XML tags.
	// Thus xml.Marshal() and the likes can be used.
	rss, _ := xml.MarshalIndent(rss2.NewRSS(channel), ``, `    `)
	fmt.Println(xml.Header + string(rss))
}
