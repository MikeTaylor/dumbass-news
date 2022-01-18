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

func MakeLogger(config LoggingConfig) *Logger {
	var logger Logger

	logger.categories = config.Categories
	logger.prefix = config.Prefix
	logger.timestamp = config.Timestamp

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
	fmt.Fprintln(os.Stderr, s+"("+cat+")", fmt.Sprintf("%v", args))
}
