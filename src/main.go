package main

import "os"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<configFile.json>")
		os.Exit(1)
	}

	var file = os.Args[1]
	var config *Config
	config, err := ReadConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read config file '%s': %v", file, err)
		os.Exit(2)
	}

	logger := MakeLogger(config.Logging)
	logger.log("config", fmt.Sprintf("%+v", config))

	server := MakeNewsServer(config, logger)
	err = server.launch(config.Listen.Host + ":" + fmt.Sprint(config.Listen.Port))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create HTTP server:", err)
		os.Exit(3)
	}
}
