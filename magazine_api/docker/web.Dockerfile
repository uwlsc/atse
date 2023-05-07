FROM golang:alpine

# Required because go requires gcc to build
RUN apk add build-base

RUN apk add inotify-tools

RUN echo $GOPATH

COPY . /magazine_api

RUN go install github.com/rubenv/sql-migrate/...@latest

WORKDIR /magazine_api

RUN go mod download

RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD sh /magazine_api/docker/run.sh