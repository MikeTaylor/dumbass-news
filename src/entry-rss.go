package main

import "fmt"
import "encoding/xml"

type channel struct {
	title string `xml:"title"`
}

type rss struct {
	XMLName  xml.Name `xml:"rss"`
	channels channel  `xml:"channel"`
}

func makeRssEntries(body []byte) ([]Entry, error) {
	fmt.Println("in makeRssEntries")

	var doc rss
	err := xml.Unmarshal([]byte("xxx"), &doc)
	if err != nil {
		return nil, fmt.Errorf("cannot parse RSS: %w", err)
	}

	fmt.Println("parse err:", err, "doc:", doc)

	return nil, nil
}

var RssEntryParser = EntryParser{
	parse: makeRssEntries,
}
