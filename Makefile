build:
#	@go build -o bin/urlshortner
	@ docker build -t url-shortner .

run: build
#	@./bin/urlshortner
	 docker run -p 8080:8080   url-shortner

test:
	@go test  -v ./...