package main

import "os"
import "fmt"

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<configFile.json>")
		os.Exit(1)
	}

	var file = os.Args[1]
	config, err := ReadConfig(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read config file", file, "-", err)
		os.Exit(2)
	}

	fmt.Printf("%+v\n", config)
}
