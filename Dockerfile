FROM golang:alpine AS builder

WORKDIR /src
COPY go.mod ./
COPY *.go ./
RUN go build -o /http-spec .

FROM alpine

LABEL maintainer="Tom Mornini <tmornini@me.com>"

RUN apk add --no-cache bash

COPY --from=builder /http-spec /usr/local/bin/http-spec

WORKDIR /
CMD ["/run-http-specs"]
