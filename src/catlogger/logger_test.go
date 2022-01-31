package catlogger

import "testing"

func assert(t *testing.T, ok bool, message string) {
	if !ok {
		t.Error(message)
	}
}

func TestLogger(t *testing.T) {
	logger := MakeLogger("foo,bar", "hello", false)
	assert(t, logger != nil, "could not make logger")
	assert(t, logger.HasCategory("foo"), "logger lacks category foo")
	assert(t, logger.HasCategory("bar"), "logger lacks category bar")
	assert(t, !logger.HasCategory("baz"), "logger has category baz")
}
