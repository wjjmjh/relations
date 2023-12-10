FROM golang:1.21 as builder

WORKDIR /go/src/relations/
COPY . .
RUN go get
RUN go build -v -ldflags "-w -s -X main.Version=1.0.0 -X main.Build=`date +%FT%T%z`" -o bin/relations-linux-amd64

FROM debian:buster-slim

MAINTAINER wjjmjh

RUN apt-get update \
    && apt-get install -y ca-certificates tzdata \
    && rm -rf /var/lib/apt/lists/*

EXPOSE 7000

WORKDIR /application

COPY --from=builder /go/src/relations/bin/relations-linux-amd64 .

ENTRYPOINT ./relations-linux-amd64
