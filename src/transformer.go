package main

type Transformer struct {
	transform func(tc TransformationConfig, entry *Entry) error
}
