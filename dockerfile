# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.20 AS build-stage
#FROM golang:1.20.10-alpine3.17 AS build-stage

ARG PORT=8080
WORKDIR /app
# COPY go.mod go.sum ./



COPY . .

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy

#RUN go get -v ./...
#RUN go install -v ./...
# COPY *.go ./
# COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /urlservice

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /
COPY /views/index.html ./views/
COPY --from=build-stage /urlservice /urlservice

EXPOSE ${PORT}



ENTRYPOINT ["./urlservice"]

