package main

import "fmt"
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
	XMLName  xml.Name `xml:"rss"`
	Channels channel  `xml:"channel"`
}

func makeRssEntries(body []byte) ([]Entry, error) {
	fmt.Println("in makeRssEntries")

	var doc rss
	err := xml.Unmarshal([]byte(body), &doc)
	if err != nil {
		return nil, fmt.Errorf("cannot parse RSS: %w", err)
	}

	fmt.Println("parse err:", err, "doc:", doc)

	return nil, nil
}

var RssEntryParser = EntryParser{
	parse: makeRssEntries,
}
