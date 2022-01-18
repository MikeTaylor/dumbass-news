package main

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
