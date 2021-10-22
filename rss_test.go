package rss2

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

type TestCase struct {
	Input    string
	Expected RSS
}

var xmlToRSSTestCases = []TestCase{
	{
		// Taken from https://cyber.harvard.edu/rss/examples/rss2sample.xml
		Input: `
		<?xml version="1.0"?>
		<rss version="2.0">
		   <channel>
		      <title>Liftoff News</title>
		      <link>http://liftoff.msfc.nasa.gov/</link>
		      <description>Liftoff to Space Exploration.</description>
		      <language>en-us</language>
		      <pubDate>Tue, 10 Jun 2003 04:00:00 GMT</pubDate>
		      <lastBuildDate>Tue, 10 Jun 2003 09:41:01 GMT</lastBuildDate>
		      <docs>http://blogs.law.harvard.edu/tech/rss</docs>
		      <generator>Weblog Editor 2.0</generator>
		      <managingEditor>editor@example.com</managingEditor>
		      <webMaster>webmaster@example.com</webMaster>
		      <item>
		         <title>Star City</title>
		         <link>http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp</link>
		         <description>How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's &lt;a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm"&gt;Star City&lt;/a&gt;.</description>
		         <pubDate>Tue, 03 Jun 2003 09:39:21 GMT</pubDate>
		         <guid>http://liftoff.msfc.nasa.gov/2003/06/03.html#item573</guid>
		      </item>
		      <item>
		         <description>Sky watchers in Europe, Asia, and parts of Alaska and Canada will experience a &lt;a href="http://science.nasa.gov/headlines/y2003/30may_solareclipse.htm"&gt;partial eclipse of the Sun&lt;/a&gt; on Saturday, May 31st.</description>
		         <pubDate>Fri, 30 May 2003 11:06:42 GMT</pubDate>
		         <guid>http://liftoff.msfc.nasa.gov/2003/05/30.html#item572</guid>
		      </item>
		      <item>
		         <title>The Engine That Does More</title>
		         <link>http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp</link>
		         <description>Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly.  The proposed VASIMR engine would do that.</description>
		         <pubDate>Tue, 27 May 2003 08:37:32 GMT</pubDate>
		         <guid>http://liftoff.msfc.nasa.gov/2003/05/27.html#item571</guid>
		      </item>
		      <item>
		         <title>Astronauts' Dirty Laundry</title>
		         <link>http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp</link>
		         <description>Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them.  Instead, astronauts have other options.</description>
		         <pubDate>Tue, 20 May 2003 08:56:02 GMT</pubDate>
		         <guid>http://liftoff.msfc.nasa.gov/2003/05/20.html#item570</guid>
		      </item>
		   </channel>
		</rss>`,
		Expected: RSS{
			XMLName: xml.Name{``, `rss`},
			Version: `2.0`,
			Channel: &Channel{
				XMLName:        xml.Name{``, `channel`},
				Title:          `Liftoff News`,
				Link:           `http://liftoff.msfc.nasa.gov/`,
				Description:    `Liftoff to Space Exploration.`,
				Language:       `en-us`,
				PubDate:        &RSSTime{time.Date(2003, 6, 10, 4, 0, 0, 0, time.FixedZone("+0000", 0))},
				LastBuildDate:  &RSSTime{time.Date(2003, 6, 10, 9, 41, 1, 0, time.FixedZone("+0000", 0))},
				Docs:           `http://blogs.law.harvard.edu/tech/rss`,
				Generator:      `Weblog Editor 2.0`,
				ManagingEditor: `editor@example.com`,
				WebMaster:      `webmaster@example.com`,
				Items: []*Item{
					{
						XMLName:     xml.Name{``, `item`},
						Title:       `Star City`,
						Link:        `http://liftoff.msfc.nasa.gov/news/2003/news-starcity.asp`,
						Description: `How do Americans get ready to work with Russians aboard the International Space Station? They take a crash course in culture, language and protocol at Russia's <a href="http://howe.iki.rssi.ru/GCTC/gctc_e.htm">Star City</a>.`,
						PubDate:     &RSSTime{time.Date(2003, 6, 3, 9, 39, 21, 0, time.FixedZone("+0000", 0))},
						GUID: &GUID{
							XMLName: xml.Name{``, `guid`},
							Value:   `http://liftoff.msfc.nasa.gov/2003/06/03.html#item573`,
						},
					},
					{
						XMLName:     xml.Name{``, `item`},
						Description: `Sky watchers in Europe, Asia, and parts of Alaska and Canada will experience a <a href="http://science.nasa.gov/headlines/y2003/30may_solareclipse.htm">partial eclipse of the Sun</a> on Saturday, May 31st.`,
						PubDate:     &RSSTime{time.Date(2003, 5, 30, 11, 6, 42, 0, time.FixedZone("+0000", 0))},
						GUID: &GUID{
							XMLName: xml.Name{``, `guid`},
							Value:   `http://liftoff.msfc.nasa.gov/2003/05/30.html#item572`,
						},
					},
					{
						XMLName:     xml.Name{``, `item`},
						Title:       `The Engine That Does More`,
						Link:        `http://liftoff.msfc.nasa.gov/news/2003/news-VASIMR.asp`,
						Description: `Before man travels to Mars, NASA hopes to design new engines that will let us fly through the Solar System more quickly.  The proposed VASIMR engine would do that.`,
						PubDate:     &RSSTime{time.Date(2003, 5, 27, 8, 37, 32, 0, time.FixedZone("+0000", 0))},
						GUID: &GUID{
							XMLName: xml.Name{``, `guid`},
							Value:   `http://liftoff.msfc.nasa.gov/2003/05/27.html#item571`,
						},
					},
					{
						XMLName:     xml.Name{``, `item`},
						Title:       `Astronauts' Dirty Laundry`,
						Link:        `http://liftoff.msfc.nasa.gov/news/2003/news-laundry.asp`,
						Description: `Compared to earlier spacecraft, the International Space Station has many luxuries, but laundry facilities are not one of them.  Instead, astronauts have other options.`,
						PubDate:     &RSSTime{time.Date(2003, 5, 20, 8, 56, 2, 0, time.FixedZone("+0000", 0))},
						GUID: &GUID{
							XMLName: xml.Name{``, `guid`},
							Value:   `http://liftoff.msfc.nasa.gov/2003/05/20.html#item570`,
						},
					},
				},
			},
		},
	},
	{
		// A constructed test case to cover many different elements
		Input: `
		<?xml version="1.0" encoding="UTF-8"?>
		<rss version="2.0">
		   <channel>
		      <title><![CDATA[Channel title with special characters: > " &]]></title>
		      <link>foo.com</link>
		      <description>Channel description</description>
		      <pubDate>Tue, 10 Jun 03 04:00:00 A</pubDate>
		      <lastBuildDate>10 Jun 2003 09:41 -0700</lastBuildDate>
		      <category domain="foo domain with escape: &gt;">Channels domain</category>
		      <cloud domain="rpc.sys.com" port="80" path="/RPC2" registerProcedure="pingMe" protocol="soap"/>
		      <image>
		          <url>image's url</url>
		          <title>image's title</title>
		          <link>image's link</link>
		          <width>80</width>
		          <height>80</height>
		          <description>image's description</description>
		      </image>
		      <textInput>
		          <title>textInput's Title</title>
		          <description>textInput's description</description>
		          <name>textInput's name</name>
		          <link>textInput's link</link>
		      </textInput>
		      <item>
		         <title>Item 1</title>
		         <pubDate>Tue, 03 Jun 93 09:39:21 MST</pubDate>
		         <guid isPermaLink="true">guid with escapes: &gt; &quot; &amp;</guid>
		         <enclosure url="enclosure's url" length="42" type="enclosure's type" />
		         <source url='source url with escapes: &gt; &quot;'>source element</source>
		      </item>
		   </channel>
		</rss>`,
		Expected: RSS{
			XMLName: xml.Name{``, `rss`},
			Version: `2.0`,
			Channel: &Channel{
				XMLName:       xml.Name{``, `channel`},
				Title:         `Channel title with special characters: > " &`,
				Link:          `foo.com`,
				Description:   `Channel description`,
				PubDate:       &RSSTime{time.Date(2003, 6, 10, 4, 0, 0, 0, time.FixedZone(`+0100`, 1*60*60))},
				LastBuildDate: &RSSTime{time.Date(2003, 6, 10, 9, 41, 0, 0, time.FixedZone(`-0700`, -7*60*60))},
				Categories: []*Category{{
					XMLName: xml.Name{``, `category`},
					Value:   `Channels domain`,
					Domain:  `foo domain with escape: >`,
				}},
				Cloud: &Cloud{
					XMLName:           xml.Name{``, `cloud`},
					Domain:            `rpc.sys.com`,
					Port:              80,
					Path:              `/RPC2`,
					RegisterProcedure: `pingMe`,
					Protocol:          `soap`,
				},
				Image: &Image{
					XMLName:     xml.Name{``, `image`},
					URL:         `image's url`,
					Title:       `image's title`,
					Link:        `image's link`,
					Width:       80,
					Height:      80,
					Description: `image's description`,
				},
				TextInput: &TextInput{
					XMLName:     xml.Name{``, `textInput`},
					Title:       `textInput's Title`,
					Description: `textInput's description`,
					Name:        `textInput's name`,
					Link:        `textInput's link`,
				},
				Items: []*Item{
					{
						XMLName: xml.Name{``, `item`},
						Title:   `Item 1`,
						PubDate: &RSSTime{time.Date(1993, 6, 3, 9, 39, 21, 0, time.FixedZone("-0700", -7*60*60))},
						GUID: &GUID{
							XMLName:     xml.Name{``, `guid`},
							Value:       `guid with escapes: > " &`,
							IsPermaLink: true,
						},
						Enclosure: &Enclosure{
							XMLName: xml.Name{``, `enclosure`},
							URL:     `enclosure's url`,
							Length:  42,
							Type:    `enclosure's type`,
						},
						Source: &Source{
							XMLName: xml.Name{``, `source`},
							Value:   `source element`,
							URL:     `source url with escapes: > "`,
						},
					},
				},
			},
		},
	},
}

func TestParseRSS(t *testing.T) {
	for _, tc := range xmlToRSSTestCases {
		var parse RSS
		if err := xml.Unmarshal([]byte(tc.Input), &parse); err != nil {
			t.Errorf(err.Error())
		}
		if diff := cmp.Diff(tc.Expected, parse); diff != "" {
			t.Errorf("RSS parsing mismatch (-want +got):\n%s", diff)
		}
	}
}

// Also trying out the constructors in this test.
func TestRenderRSS(t *testing.T) {
	item1, err := NewItem(`Item 1`, ``)
	if err != nil {
		t.Errorf(err.Error())
	}
	item1.PubDate = &RSSTime{time.Date(1993, 6, 3, 9, 39, 21, 0,
		time.FixedZone("-0700", -7*60*60))}
	if item1.GUID, err = NewGUID(`guid with escapes: > " &`); err != nil {
		t.Errorf(err.Error())
	}
	item1.GUID.IsPermaLink = true
	item1.Enclosure, err = NewEnclosure(`enclosure's url`, 42, `enclosure's type`)
	if err != nil {
		t.Errorf(err.Error())
	}
	item1.Source, err = NewSource(`source element`, `source url with escapes: > "`)
	if err != nil {
		t.Errorf(err.Error())
	}

	item2, err := NewItem(`Item 2`, ``)
	if err != nil {
		t.Errorf(err.Error())
	}

	channel, err := NewChannel(
		`Channel title with escapes: > " &`, `foo.com`, `Channel description`)
	if err != nil {
		t.Errorf(err.Error())
	}
	channel.PubDate = &RSSTime{time.Date(2003, 6, 10, 4, 0, 0, 0,
		time.FixedZone(`+0100`, 1*60*60))}
	channel.LastBuildDate = &RSSTime{time.Date(2003, 6, 10, 9, 41, 0, 0,
		time.FixedZone(`-0700`, -7*60*60))}
	category, err := NewCategory(`Categorie's domain`)
	if err != nil {
		t.Errorf(err.Error())
	}
	category.Domain = `foo domain with escape: >`
	channel.Categories = []*Category{category}
	channel.Cloud, err = NewCloud(`rpc.sys.com`, 80, `/RPC2`, `pingMe`, `soap`)
	if err != nil {
		t.Errorf(err.Error())
	}
	channel.Image, err = NewImage(`image's url`, `image's title`, `image's link`)
	if err != nil {
		t.Errorf(err.Error())
	}
	channel.Image.Width = 80
	channel.Image.Height = 80
	channel.Image.Description = `image's description`
	channel.TextInput, err = NewTextInput(`textInput's Title`,
		`textInput's description`, `textInput's name`, `textInput's link`)
	if err != nil {
		t.Errorf(err.Error())
	}
	if channel.SkipHours, err = NewSkipHours([]int{0, 1, 2, 3, 4}); err != nil {
		t.Errorf(err.Error())
	}
	channel.SkipDays = NewSkipDays([]time.Weekday{time.Sunday})
	channel.Items = []*Item{item1, item2}

	out, err := xml.MarshalIndent(NewRSS(channel), "\t", `    `)
	if err != nil {
		t.Errorf(err.Error())
	}

	expected := `	<rss version="2.0">
	    <channel>
	        <title>Channel title with escapes: &gt; &#34; &amp;</title>
	        <link>foo.com</link>
	        <description>Channel description</description>
	        <pubDate>10 Jun 2003 04:00:00 +0100</pubDate>
	        <lastBuildDate>10 Jun 2003 09:41:00 -0700</lastBuildDate>
	        <category domain="foo domain with escape: &gt;">Categorie&#39;s domain</category>
	        <cloud domain="rpc.sys.com" port="80" path="/RPC2" registerProcedure="pingMe" protocol="soap"></cloud>
	        <image>
	            <url>image&#39;s url</url>
	            <title>image&#39;s title</title>
	            <link>image&#39;s link</link>
	            <width>80</width>
	            <height>80</height>
	            <description>image&#39;s description</description>
	        </image>
	        <textInput>
	            <title>textInput&#39;s Title</title>
	            <description>textInput&#39;s description</description>
	            <name>textInput&#39;s name</name>
	            <link>textInput&#39;s link</link>
	        </textInput>
	        <skipHours>
	            <hour>0</hour>
	            <hour>1</hour>
	            <hour>2</hour>
	            <hour>3</hour>
	            <hour>4</hour>
	        </skipHours>
	        <skipDays>
	            <day>Sunday</day>
	        </skipDays>
	        <item>
	            <title>Item 1</title>
	            <enclosure url="enclosure&#39;s url" length="42" type="enclosure&#39;s type"></enclosure>
	            <guid isPermaLink="true">guid with escapes: &gt; &#34; &amp;</guid>
	            <pubDate>03 Jun 1993 09:39:21 -0700</pubDate>
	            <source url="source url with escapes: &gt; &#34;">source element</source>
	        </item>
	        <item>
	            <title>Item 2</title>
	        </item>
	    </channel>
	</rss>`

	if diff := cmp.Diff(expected, string(out)); diff != "" {
		t.Errorf("RSS parsing mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkParseRSS(b *testing.B) {
	var parse RSS
	for i := 0; i < b.N; i++ {
		xml.Unmarshal([]byte(xmlToRSSTestCases[0].Input), &parse)
	}
}
