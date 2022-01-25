package main

type Entry struct {
	Headline string
	Link     string
}

type EntryParser struct {
	parse func(cc ChannelConfig, body []byte) ([]Entry, error)
}
