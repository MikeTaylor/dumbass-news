package main

import "net/http"
import "fmt"
import "strings"
import "time"
import "io/ioutil"

func renderHTML(w http.ResponseWriter, server *NewsServer, channel string, transformation string, entries []Entry) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "<h1>channel '%s'</h1>\n", channel)
	fmt.Fprintf(w, "<p>(after transformation '%s')</p>\n", transformation)
	fmt.Fprintf(w, "<ul>\n")
	for i := 0; i < len(entries); i++ {
		fmt.Fprintf(w, "<li><a href=\"%s\">%s</a></li>\n", entries[i].Link, entries[i].Headline)
	}
	fmt.Fprintf(w, "</ul>\n")
}

func getData(server *NewsServer, channel string) ([]Entry, *httpError) {
	channelConfig, ok := server.config.Channels[channel]
	if !ok {
		return nil, MakeHttpError(http.StatusBadRequest, fmt.Sprintln("unknown channel:", channel))
	}
	server.logger.log("config", fmt.Sprintf("channel '%s': %+v", channel, channelConfig))

	ctype := channelConfig.ChannelType
	var parser EntryParser
	switch ctype {
	case "rss":
		parser = RssEntryParser
	// more cases here
	default:
		return nil, MakeHttpError(http.StatusBadRequest, fmt.Sprintln("unsupported channel-type:", ctype))
	}

	resp, err := server.client.Get(channelConfig.Url)
	if err != nil {
		return nil, MakeHttpError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch %s: %v", channelConfig.Url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, MakeHttpError(http.StatusInternalServerError, fmt.Sprintf("fetch %s failed with status %s", channelConfig.Url, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	server.logger.log("body", fmt.Sprintf("%s", body))

	var entries []Entry
	entries, err = parser.parse(body)
	if err != nil {
		return nil, MakeHttpError(http.StatusInternalServerError, fmt.Sprintf("parsing source %s failed: %v", channelConfig.Url, err))
	}

	return entries, nil
}

func showChannel(w http.ResponseWriter, server *NewsServer, channel string, transformation string) {
	entries, err := getData(server, channel)
	if entries == nil {
		w.WriteHeader(err.status)
		fmt.Fprint(w, err.message)
		return
	}
	// XXX transform data

	renderHTML(w, server, channel, transformation, entries)
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
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir("..")))

	var server = NewsServer{
		config: config,
		logger: logger,
		server: http.Server{
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
		client: http.Client{
			Timeout:   10 * time.Second,
			Transport: tr,
		},
	}

	// XXX I would prefer this to be registered only to server.server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handler(w, r, &server) })

	return &server
}

func (server *NewsServer) launch(hostspec string) error {
	server.server.Addr = hostspec
	server.logger.log("listen", "listening on", hostspec)
	err := server.server.ListenAndServe()
	server.logger.log("listen", "finished listening on", hostspec)
	return err
}
