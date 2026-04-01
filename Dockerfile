FROM golang:alpine
MAINTAINER Tom Mornini <tmornini@me.com>

RUN apk update   && \
    apk add bash && \
    apk add git

COPY * /go/src/github.com/tmornini/http-spec/

RUN cd /go/src/github.com/tmornini/http-spec && \
    go install .

WORKDIR /

CMD ["/run-http-specs"]
