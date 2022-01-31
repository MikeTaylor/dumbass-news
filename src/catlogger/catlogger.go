// catlogger is derived from the JavaScript package
// categorical-logger, which can be found at
// https://github.com/openlibraryenvironment/categorical-logger
// The present version of this library falls short of its ancestor in
// two respects:
//  1. It does not support logging functional arguments
//  2. It does not have the the setter methods
//
package catlogger

import "strings"
import "time"
import "fmt"
import "os"

// Logger is an opaque structure created by MakeLogger and which
// encapsulates the configuration passed into that function. Typically
// a program will make just one of these, and pass it around as
// necessary.
//
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

// MakeLogger creates a catlogger object on which the Log
// method may subsequently be called. The logger is configured by
// three parameters:
//
// categories: a string containing zero or more comma-separated
// logging categories, such as "init,config,url", for which the logger
// should emit messages. There is no predefined list of categegories,
// and no hierarchy between categories -- unlike the FATAL, ERROR,
// WARN, INFO, DEBUG hierarchy in log4j and similar
// libraries. Applications can use whatever categories they wish.
//
// prefix: a short string or nil. If the former, then it is included
// at the start of each log message.
//
// timestamp: a boolean indicating whether or not a timestamp is to be
// included in each log message
//
func MakeLogger(categories string, prefix string, timestamp bool) *Logger {
	var logger Logger

	logger.categories = getCategories(categories)
	logger.prefix = prefix
	logger.timestamp = timestamp

	return &logger
}

// HasCategory returns true is the logger has been configured to
// include the named category and false if not.
//
func (l *Logger) HasCategory(cat string) bool {
	var xlist string = "," + l.categories + ","
	var xcat string = "," + cat + ","
	return strings.Contains(xlist, xcat)
}

// Log does nothing if the named category is not among those that the
// logger is configured to use. But if the logged is configured to
// include the category, then it logs a message to standard error,
// consisting of the logger prefix (if defined), a timestamp (if
// configured), and all the supplied strings, separated by spaces and
// terminated by a newline.
//
func (l *Logger) Log(cat string, args ...string) {
	if !l.HasCategory(cat) {
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
