package main

import "testing"
import "fmt"
import "os"
import "time"
import "net/http"
import "io/ioutil"
import "regexp"

func TestDumbassNews(t *testing.T) {
	// For now we duplicate code from main.go

	file := "../etc/config.json"
	var cfg *config
	cfg, err := readConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot read config file '%s': %v\n", file, err)
		os.Exit(2)
	}

	cl := cfg.Logging
	logger := MakeLogger("listen", cl.Prefix, cl.Timestamp)
	logger.log("config", fmt.Sprintf("%+v", cfg))

	server := MakeNewsServer(cfg, logger, "..")
	go func() {
		err = server.launch(cfg.Listen.Host + ":12369")
	}()

	// Allow half a second for the server to start. This is ugly
	time.Sleep(time.Second / 2)
	runTests(t, server.client)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot create HTTP server:", err)
		os.Exit(3)
	}
}

func runTests(t *testing.T, client http.Client) {
	data := []struct {
		name   string
		path   string
		status int
		re     string
	}{
		{"home", "", 200, "Example pages"},
		{"short path", "foo", 404, ""},
		{"long path", "foo/bar/baz", 404, ""},
		{"bad channel", "badchannel/bar", 400, "unknown channel"},
		{"bad channel-type", "badtype/bar", 400, "unsupported channel-type"},
		{"bad RSS host", "nohost/disemvowel", 500, "no such host"},
		{"bad RSS file", "nofile/disemvowel", 500, "404 Not Found"},
		{"local RSS", "static/disemvowel", 200, "Crmnl prsctn"},
		{"remote RSS", "bbc/disemvowel", 200, "https://www.bbc.co.uk/news/"},     // Links from this feed
		{"Hacker News RSS with insert", "hackernews/dumbass", 200, "hackernews"}, // Links could be to anywhere
		{"SV-POW! RSS with insert", "svpow/dumbass", 200, "https://svpow.com/.*dumbass"},
		{"static CSS file", "htdocs/style.css", 200, "text-decoration"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			url := "http://localhost:12369/" + d.path
			resp, err := client.Get(url)
			if err != nil {
				t.Errorf("cannot fetch %s: %v", url, err)
				return
			}
			defer resp.Body.Close()
			if resp.StatusCode != d.status {
				t.Errorf("fetch %s had status %s (expected %d)", url, resp.Status, d.status)
				// Do not return; attempt the remaining checks
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read body %s: %v", url, err)
				return
			}
			matched, err := regexp.Match(d.re, body)
			if err != nil {
				t.Errorf("cannot match body of %s against regexp /%s/: %v", url, d.re, err)
				return
			}
			if !matched {
				t.Errorf("body of %s does not match regexp /%s/: body = %s", url, d.re, body)
			}
		})
	}
}
