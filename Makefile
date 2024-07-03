# .PHONY: build clean deploy gomodgen

build: clean
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/main cmd/app/main.go

deploy_prod: build
	serverless deploy --stage prod

dev:
	sam local start-api

clean:
	go clean
	rm -rf ./bin

