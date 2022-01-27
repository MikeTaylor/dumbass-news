package main

import "net/http"
import "fmt"
import "strings"
import "time"
import "io/ioutil"

func showHome(w http.ResponseWriter, server *NewsServer) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<link rel="stylesheet" type="text/css" href="/htdocs/style.css"/>`)
	fmt.Fprintln(w, "<h1>Example pages</h2>")
	fmt.Fprintln(w, `<ul>
<li><a href="http://localhost:12368/">200 (this page)</a></li>
<li><a href="http://localhost:12368/foo">404 (no transformation component)</a></li>
<li><a href="http://localhost:12368/foo/disemvowel/baz">404 (extra path component)</a></li>
<li><a href="http://localhost:12368/foo/disemvowel">400 (bad channel)</a></li>
<li><a href="http://localhost:12368/badtype/disemvowel">400 (channel of bad type)</a></li>
<li><a href="http://localhost:12368/nohost/disemvowel">500 (RSS channel with bad host)</a></li>
<li><a href="http://localhost:12368/nofile/disemvowel">500 (RSS channel with bad file)</a></li>
<li><a href="http://localhost:12368/bbc/disemvowel">200 (RSS channel working)</a></li>
<li><a href="http://localhost:12368/static/disemvowel">200 (RSS channel working from static file)</a></li>
<li><a href="http://localhost:12368/bbc/dumbass">200 (BBC RSS channel with dumbass transformation)</a></li>
<li><a href="http://localhost:12368/hackernews/dumbass">200 (Hacker News RSS channel with dumbass transformation)</a></li>
<li><a href="http://localhost:12368/svpow/dumbass">200 (SV-POW! RSS channel with dumbass transformation)</a></li>`)
	fmt.Fprintln(w, "</ul>")
}

func renderHTML(w http.ResponseWriter, server *NewsServer, channel string, transformation string, entries []Entry) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, `<link rel="stylesheet" type="text/css" href="/htdocs/style.css"/>`)
	fmt.Fprintf(w, `<h1>
  <div class="homelink"><a href="/">home</a></div>
  <div class="channel">%s</div>
  <div class="transformation">%s</div>
</h1>
`, channel, transformation)
	fmt.Fprintf(w, "<ul>\n")
	for i := 0; i < len(entries); i++ {
		fmt.Fprintf(w, "  <li><a href=\"%s\">%s</a></li>\n", entries[i].Link, entries[i].Headline)
	}
	fmt.Fprintf(w, "</ul>\n")
}

func getData(server *NewsServer, cc channelConfig) ([]Entry, *httpError) {
	ctype := cc.ChannelType
	var parser EntryParser
	switch ctype {
	case "rss":
		parser = RssEntryParser
	// more cases here
	default:
		return nil, MakeHttpError(http.StatusBadRequest, fmt.Sprintln("unsupported channel-type:", ctype))
	}

	resp, err := server.client.Get(cc.Url)
	if err != nil {
		return nil, MakeHttpError(http.StatusInternalServerError, fmt.Sprintf("cannot fetch %s: %v", cc.Url, err))
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, MakeHttpError(http.StatusInternalServerError, fmt.Sprintf("fetch %s failed with status %s", cc.Url, resp.Status))
	}

	body, err := ioutil.ReadAll(resp.Body)
	server.logger.log("body", fmt.Sprintf("%s", body))

	var entries []Entry
	entries, err = parser.parse(cc, body)
	if err != nil {
		return nil, MakeHttpError(http.StatusInternalServerError, fmt.Sprintf("parsing source %s failed: %v", cc.Url, err))
	}

	return entries, nil
}

func showChannel(w http.ResponseWriter, server *NewsServer, channel string, transformation string) {
	channelConfig, ok := server.config.Channels[channel]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "unknown channel:", channel)
		return
	}
	server.logger.log("config", fmt.Sprintf("channel '%s': %+v", channel, channelConfig))

	entries, httpErr := getData(server, channelConfig)
	if httpErr != nil {
		w.WriteHeader(httpErr.status)
		fmt.Fprintln(w, httpErr.message)
		return
	}

	tc, ok := server.config.Transformations[transformation]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "unknown transformation:", transformation)
		return
	}
	server.logger.log("config", fmt.Sprintf("transformation '%s': %+v", transformation, tc))
	err := transformData(server, tc, entries)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err.Error())
		return
	}

	renderHTML(w, server, channel, transformation, entries)
}

func handler(w http.ResponseWriter, req *http.Request, server *NewsServer) {
	raw := strings.Split(req.URL.Path, "/")
	chunks := raw[1:] // Skip initial empty component
	server.logger.log("path", strings.Join(chunks, "/"))

	if len(chunks) == 1 && chunks[0] == "" {
		// Home page
		showHome(w, server)
	} else if len(chunks) == 2 {
		// Transformed channel
		showChannel(w, server, chunks[0], chunks[1])
	} else {
		// Unrecognized
		w.WriteHeader(http.StatusNotFound)
	}
}

type NewsServer struct {
	config *config
	logger *Logger
	server http.Server
	client http.Client
}

func MakeNewsServer(cfg *config, logger *Logger) *NewsServer {
	tr := &http.Transport{}
	tr.RegisterProtocol("file", http.NewFileTransport(http.Dir(".")))

	var server = NewsServer{
		config: cfg,
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

	fs := http.FileServer(http.Dir("./htdocs"))
	http.Handle("/htdocs/", http.StripPrefix("/htdocs/", fs))

	return &server
}

func (server *NewsServer) launch(hostspec string) error {
	server.server.Addr = hostspec
	server.logger.log("listen", "listening on", hostspec)
	err := server.server.ListenAndServe()
	server.logger.log("listen", "finished listening on", hostspec)
	return err
}
