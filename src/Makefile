dumbass-news: main.go config.go logger.go server.go http-error.go entry.go entry-rss.go transformer.go transformer-disemvowel.go transformer-insert.go
	go build -o dumbass-news ./...

run: dumbass-news
	cd .. && src/dumbass-news etc/config.json

lint:
	go vet ./...

test:
	go test -v ./...

clean:
	rm -f dumbass-news

