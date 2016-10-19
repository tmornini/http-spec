package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sync"
)

type context struct {
	HTTPS            bool
	HostName         string
	URIScheme        string
	Pathname         string
	WaitGroup        *sync.WaitGroup
	Scanner          *bufio.Scanner
	Request          *request
	ExpectedResponse *response
	HTTPClient       *http.Client
	HTTPResponse     *http.Response
	ActualResponse   *response
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
