package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
)

type context struct {
	Pathname         string
	HTTPS            bool
	HostName         string
	URIScheme        string
	Scanner          *bufio.Scanner
	Request          *request
	ExpectedResponse *response
	HTTPClient       *http.Client
	HTTPResponse     *http.Response
	NetConnection    net.Conn
	ActualResponse   *response
	ActualLine       string
	Substitutions    map[string]string
}

func (context *context) log(tag string) {
	if logFunctions {
		fmt.Println(tag)
	}

	if logContext {
		fmt.Printf("%#v\n", context)
	}
}
