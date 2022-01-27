// logger is derived from the JavaScript package categorical-logger,
// which can be found at
// https://github.com/openlibraryenvironment/categorical-logger
// The present version of logger.go falls short of its ancestor in the
// following respects:
// 1. It does not support logging functional arguments
// 2. It does not have the the setter methods
// 3. It is not documented
//
package main

import "strings"
import "time"
import "fmt"
import "os"

type Logger struct {
	categories string
	prefix     string
	timestamp  bool
}

func getCategories(fallback string) string {
	res := os.Getenv("LOGGING_CATEGORIES")
	if res != "" {
		return res
	}
	res = os.Getenv("LOGCAT")
	if res != "" {
		return res
	}
	return fallback
}

func MakeLogger(categories string, prefix string, timestamp bool) *Logger {
	var logger Logger

	logger.categories = getCategories(categories)
	logger.prefix = prefix
	logger.timestamp = timestamp

	return &logger
}

func (l *Logger) hasCategory(cat string) bool {
	var xlist string = "," + l.categories + ","
	var xcat string = "," + cat + ","
	return strings.Contains(xlist, xcat)
}

func (l *Logger) log(cat string, args ...string) {
	if !l.hasCategory(cat) {
		return
	}

	var s = ""
	if l.prefix != "" {
		s = l.prefix + " "
	}
	if l.timestamp {
		s += time.Now().Format(time.RFC3339) + " "
	}
	fmt.Fprintln(os.Stderr, s+"("+cat+")", strings.Join(args, " "))
}
