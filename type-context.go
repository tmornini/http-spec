package main

import (
	"fmt"
	"math/big"
	"net/http"
	"sync"
	"time"
)

type context struct {
	Err                   error
	File                  *file
	Hostname              string
	HTTPClient            *http.Client
	HTTPResponse          *http.Response
	HTTPRetryDelay        time.Duration
	ID                    *big.Int
	LogContext            bool
	LogFunctions          bool
	Matchers              map[string]string
	MaxHTTPAttempts       int
	Pathname              string
	Pathnames             []string
	ResultGathererChannel chan context
	SkipTLSVerification   bool
	SpecTriplet           *specTriplet
	Stage                 string
	StartedAt             time.Time
	Substitutions         map[string]string
	Scheme                string
	WaitGroup             *sync.WaitGroup
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
