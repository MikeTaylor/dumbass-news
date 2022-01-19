package main

type Entry struct {
	Headline string
	Link     string
}

// XXX This feels like it wants to be an interface
type EntryParser struct {
	parse func(body []byte) []Entry
}
