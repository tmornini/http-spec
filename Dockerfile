FROM golang:alpine
MAINTAINER Tom Mornini <tom@subledger.com>

RUN apk update   && \
    apk add bash

COPY http-spec /

ENV PATH="/:${PATH}"
