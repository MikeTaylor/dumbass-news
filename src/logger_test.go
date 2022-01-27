package main

import "testing"

func assert(t *testing.T, ok bool, message string) {
	if !ok {
		t.Error(message)
	}
}

func TestLogger(t *testing.T) {
	logger := MakeLogger("foo,bar", "hello", false)
	assert(t, logger != nil, "could not make logger")
	assert(t, logger.hasCategory("foo"), "logger lacks category foo")
	assert(t, logger.hasCategory("bar"), "logger lacks category bar")
	assert(t, !logger.hasCategory("baz"), "logger has category baz")
}
