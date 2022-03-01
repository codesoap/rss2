[![GoDoc](https://godoc.org/github.com/codesoap/rss2?status.svg)](https://godoc.org/github.com/codesoap/rss2)

A complete and strictly
[standard](https://cyber.harvard.edu/rss/rss.html) conforming library
for parsing and rendering RSS 2.0 feeds.

Simple example which renders a feed:

```go
package main

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/codesoap/rss2"
)

func main() {
	item, _ := rss2.NewItem(`Stonehenge finally understood!`, ``)
	item.Link = `https://willies-wilts.news/stonehenge-understood`
	item.PubDate = &rss2.RSSTime{time.Date(2022, 2, 3, 9, 39, 21, 0, time.UTC)}

	channelTitle := `Willie's Wiltshire News`
	channelLink := `https://willies-wilts.news`
	channelDesc := `Willie's latest news regarding Wiltshire`
	channel, _ := rss2.NewChannel(channelTitle, channelLink, channelDesc)
	channel.Items = []*rss2.Item{item}

	rss, _ := xml.MarshalIndent(rss2.NewRSS(channel), ``, `    `)
	fmt.Println(xml.Header + string(rss))
}
```

Find more examples and documentation at
[https://pkg.go.dev/github.com/codesoap/rss2](https://pkg.go.dev/github.com/codesoap/rss2).
