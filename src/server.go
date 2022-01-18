package main

import "net/http"
import "time"

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
	return &server
}

func (server *HTTPServer) ListenAndServe(hostspec string) error {
	server.server.Addr = hostspec
	server.logger.log("listen", "listening on", hostspec)
	error := server.server.ListenAndServe()
	server.logger.log("listen", "finished listening on", hostspec)
	return error
}
