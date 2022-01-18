package main

import "os"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<configFile.json>")
		os.Exit(1)
	}

	var file = os.Args[1]
	var config *Config;
	config, err := ReadConfig(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read config:", err)
		os.Exit(2)
	}

	server, err := MakeHTTPServer(config)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create HTTP server:", err)
		os.Exit(3)
	}

	fmt.Printf("%+v\n", config)
	server.ListenAndServe(":8090", nil)
}
