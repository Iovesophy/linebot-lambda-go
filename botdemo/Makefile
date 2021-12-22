.PHONY: build clean deploy

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/bot bot/main.go bot/line.go

clean:
	rm -rf ./bin

deploy:
	sls deploy --verbose
