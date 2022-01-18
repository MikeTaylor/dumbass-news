dumbass-news: main.go config.go logger.go server.go
	go build -o dumbass-news *.go

run: dumbass-news
	./dumbass-news ../etc/config.json

lint:
	go vet *.go

clean:
	rm -f dumbass-news

