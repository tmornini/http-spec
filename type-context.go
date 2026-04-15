package main

import (
	"fmt"
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
	HTTPTimeout           time.Duration
	ID                    string
	LogContext            bool
	LogFunctions          bool
	Matchers              map[string]string
	MatchersPath          string
	MaxHTTPAttempts       int
	RetryStatusCodes      []int
	Pathname              string
	Pathnames             []string
	ResultGathererChannel chan context
	ShowSubstitutions     bool
	SkipTLSVerification   bool
	SpecTriplet           *specTriplet
	Stage                 string
	StartedAt             time.Time
	Substitutions         map[string]string
	Scheme                string
	WaitGroup             *sync.WaitGroup
}

func (context *context) isRetryStatusCode() bool {
	for _, code := range context.RetryStatusCodes {
		if context.HTTPResponse.StatusCode == code {
			return true
		}
	}

	return false
}

func (context *context) enterStage(stage string) {
	context.Stage = stage

	if context.LogFunctions {
		fmt.Println(stage)
	}

	if context.LogContext {
		fmt.Printf("%#v\n", context)
	}
}
