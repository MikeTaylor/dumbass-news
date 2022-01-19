package main

import "net/http"
import "fmt"
import "strings"
import "time"

func showChannel(w http.ResponseWriter, server *NewsServer, channel string, transformation string) {
	channelConfig, ok := server.config.Channels[channel]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "unknown channel:", channel)
		return
	}
	server.logger.log("config", fmt.Sprintf("channel '%s': %+v", channel, channelConfig))

	ctype := channelConfig.ChannelType
	if ctype != "rss" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "unsupported channel-type:", ctype)
		return
	}
		
	fmt.Fprintln(w, "channel:", channel, "- transformation:", transformation)
}

func handler(w http.ResponseWriter, req *http.Request, server *NewsServer) {
	raw := strings.Split(req.URL.Path, "/")
	chunks := raw[1:] // Skip initial empty component
	server.logger.log("path", strings.Join(chunks, "/"))

	if len(chunks) == 1 && chunks[0] == "" {
		// Home page
		fmt.Fprintln(w, `<a href="/bbc/dumbass">Example</a>`)
	} else if len(chunks) == 2 {
		// Transformed channel
		showChannel(w, server, chunks[0], chunks[1])
	} else {
		// Unrecognized
		w.WriteHeader(http.StatusNotFound)
	}
}

type NewsServer struct {
	config *Config
	logger *Logger
	server http.Server
	client http.Client
}

func MakeNewsServer(config *Config, logger *Logger) *NewsServer {
	var server = NewsServer{
		config: config,
		logger: logger,
		server: http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	// XXX I would prefer this to be registered only to server.server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handler(w, r, &server) })

	return &server
}

func (server *NewsServer) launch(hostspec string) error {
	server.server.Addr = hostspec
	server.logger.log("listen", "listening on", hostspec)
	error := server.server.ListenAndServe()
	server.logger.log("listen", "finished listening on", hostspec)
	return error
}
