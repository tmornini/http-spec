package main

import (
	"fmt"
	"net/http"
	"sync"
)

type context struct {
	HTTPS                    bool
	HostName                 string
	URIScheme                string
	File                     *file
	WaitGroup                *sync.WaitGroup
	Request                  *request
	ExpectedResponse         *response
	HTTPClient               *http.Client
	HTTPResponse             *http.Response
	ActualResponse           *response
	ActualResponseLineNumber int
	Substitutions            map[string]string
}

func (context *context) log(tag string) {
	if logFunctions {
		fmt.Println(tag)
	}

	if logContext {
		fmt.Printf("%#v\n", context)
	}
}
