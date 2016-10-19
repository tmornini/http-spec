package main

import (
	"fmt"
	"net/http"
	"sync"
)

type context struct {
	HTTPS              bool
	URIScheme          string
	HostName           string
	Pathnames          []string
	Pathname           string
	WaitGroup          *sync.WaitGroup
	File               *file
	HTTPClient         *http.Client
	SpecTriplet        *specTriplet
	Substitutions      map[string]string
	HTTPResponse       *http.Response
	ResponseLineNumber int
}

func (context *context) log(tag string) {
	if logFunctions {
		fmt.Println(tag)
	}

	if logContext {
		fmt.Printf("%#v\n", context)
	}
}
