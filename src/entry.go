package main

type Entry struct {
	Headline string
	Link     string
}

type EntryParser struct {
	parse func(body []byte) ([]Entry, error)
}
