package main

import "fmt"
import "errors"
import "encoding/xml"

type item struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
}

type channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	Items       []item `xml:"item"`
}

type rss struct {
	XMLName  xml.Name  `xml:"rss"`
	Channels []channel `xml:"channel"`
}

func makeRssEntries(cc channelConfig, body []byte) ([]Entry, error) {
	var doc rss
	err := xml.Unmarshal([]byte(body), &doc)
	if err != nil {
		return nil, fmt.Errorf("cannot parse RSS: %w", err)
	}

	// XXX Not sure at the moment what to do with top-level
	// metadata. For the time being I will just discard it, along
	// with all channels but the first; and merge the item title
	// and description into a single field.

	if doc.Channels == nil || len(doc.Channels) == 0 {
		return nil, errors.New("no channels in RSS feed")
	}
	ch := doc.Channels[0]
	res := make([]Entry, len(ch.Items))
	for i := 0; i < len(ch.Items); i++ {
		item := ch.Items[i]
		if cc.Render == "title" {
			res[i].Headline = item.Title
		} else if cc.Render == "description" {
			res[i].Headline = item.Description
		} else {
			res[i].Headline = item.Title + " -- " + item.Description
		}
		res[i].Link = item.Link
	}

	return res, nil
}

var RssEntryParser = EntryParser{
	parse: makeRssEntries,
}
