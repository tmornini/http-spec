package main

import (
	"fmt"
	"net/http"
	"sync"
)

type context struct {
	LogFunctions          bool
	LogContext            bool
	Stage                 string
	HTTPS                 bool
	URIScheme             string
	HostName              string
	Pathnames             []string
	Pathname              string
	WaitGroup             *sync.WaitGroup
	ResultGathererChannel chan context
	File                  *file
	Substitutions         map[string]string
	HTTPClient            *http.Client
	SpecTriplet           *specTriplet
	HTTPResponse          *http.Response
	Err                   error
}

func (context *context) log(stage string) {
	context.Stage = stage

	if context.LogFunctions {
		fmt.Println(stage)
	}

	if context.LogContext {
		fmt.Printf("%#v\n", context)
	}
}
