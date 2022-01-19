package main

import "net/http"
import "fmt"
import "strings"
import "time"

type HTTPServer struct {
	config *Config
	logger *Logger
	server http.Server
}

func handler(w http.ResponseWriter, req *http.Request) {
	raw := strings.Split(req.URL.Path, "/")
	chunks := raw[1:] // Skip initial empty component

	if len(chunks) == 1 && chunks[0] == "" {
		// Home page
		fmt.Fprintln(w, `<a href="/bbc/dumbass">Example</a>`)
	} else if len(chunks) == 2 {
		// Transformed channel
		fmt.Fprintln(w, "channel", chunks[0], "transformation", chunks[1])
	} else {
		// Unrecognized
		w.WriteHeader(http.StatusNotFound)
	}
}

func MakeHTTPServer(config *Config, logger *Logger) *HTTPServer {
	var server = HTTPServer{
		config: config,
		logger: logger,
		server: http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}

	// XXX I would prefer this to be registered only to server.server
	http.HandleFunc("/", handler)

	return &server
}

func (server *HTTPServer) launch(hostspec string) error {
	server.server.Addr = hostspec
	server.logger.log("listen", "listening on", hostspec)
	error := server.server.ListenAndServe()
	server.logger.log("listen", "finished listening on", hostspec)
	return error
}
