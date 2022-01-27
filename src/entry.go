package main

type Entry struct {
	Headline string
	Link     string
}

type EntryParser struct {
	parse func(cc channelConfig, body []byte) ([]Entry, error)
}
