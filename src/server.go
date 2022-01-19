package main

import "net/http"
import "fmt"
import "strings"
import "time"

func showChannel(w http.ResponseWriter, config *Config, logger *Logger, channel string, transformation string) {
	fmt.Fprintln(w, "channel:", channel, "- transformation:", transformation)
}

func handler(w http.ResponseWriter, req *http.Request, config *Config, logger *Logger) {
	raw := strings.Split(req.URL.Path, "/")
	chunks := raw[1:] // Skip initial empty component
	logger.log("path", strings.Join(chunks, "/"))

	if len(chunks) == 1 && chunks[0] == "" {
		// Home page
		fmt.Fprintln(w, `<a href="/bbc/dumbass">Example</a>`)
	} else if len(chunks) == 2 {
		// Transformed channel
		showChannel(w, config, logger, chunks[0], chunks[1])
	} else {
		// Unrecognized
		w.WriteHeader(http.StatusNotFound)
	}
}

type HTTPServer struct {
	config *Config
	logger *Logger
	server http.Server
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handler(w, r, config, logger) })

	return &server
}

func (server *HTTPServer) launch(hostspec string) error {
	server.server.Addr = hostspec
	server.logger.log("listen", "listening on", hostspec)
	error := server.server.ListenAndServe()
	server.logger.log("listen", "finished listening on", hostspec)
	return error
}
