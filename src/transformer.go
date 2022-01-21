package main

type Transformer struct {
	transform func(entry *Entry) error
}
