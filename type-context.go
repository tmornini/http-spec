package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type context struct {
	LogFunctions          bool
	LogContext            bool
	URLPrefix             string
	Pathnames             []string
	Pathname              string
	WaitGroup             *sync.WaitGroup
	ResultGathererChannel chan context
	Stage                 string
	File                  *file
	Substitutions         map[string]string
	HTTPClient            *http.Client
	SpecTriplet           *specTriplet
	StartedAt             time.Time
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
