FROM golang:alpine
MAINTAINER Tom Mornini <tmornini@incentivenetworks.com>

RUN apk update   && \
    apk add bash && \
    apk add git

RUN (                                                        \
      mkdir -p /go/src/github.com/tmornini                && \
      cd /go/src/github.com/tmornini                      && \
      git clone https://github.com/tmornini/http-spec.git && \
      cd http-spec                                        && \
      go install .                                           \
    )

WORKDIR /

CMD ["/run-http-specs"]
