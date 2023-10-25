build:
	@go build -o bin/urlshortner

run: build
	@./bin/urlshortner

test:
	@go test  -v ./...