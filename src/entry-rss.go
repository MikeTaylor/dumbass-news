package main

import "fmt"

func makeRssEntries(body []byte) []Entry {
	fmt.Println("in makeRssEntries")
	return nil
}

var RssEntryParser = EntryParser{
	parse: makeRssEntries,
}
