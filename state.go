package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var regexpIdentifier = "⧆"
var substitutionIdentifier = "⧈"

type state struct {
	Error              error
	File               *file
	HTTPClient         *http.Client
	HTTPResponse       *http.Response
	HTTPRetryDelay     time.Duration
	ID                 string
	LogState           bool
	LogTags            bool
	MaxHTTPAttempts    int
	HTSFFilePath       string
	HTSFFilePaths      []string
	ResponseLineNumber int
	SpecTriplet        *specTriplet
	StartedAt          time.Time
	EndedAt            time.Time
	Result             *result
	Substitutions      map[string]string
	WaitGroup          *sync.WaitGroup
}

func NewState(skipTLSVerification ...bool) *state {
	s := &state{
		HTTPRetryDelay:  time.Second,
		LogState:        false,
		LogTags:         false,
		MaxHTTPAttempts: 180,
		StartedAt:       time.Now(),
		Result:          &result{},
		Substitutions:   make(map[string]string),
		WaitGroup:       &sync.WaitGroup{},
	}

	skip := false

	if len(skipTLSVerification) > 0 {
		skip = skipTLSVerification[0]
	}

	if skip {
		s.HTTPClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	} else {
		s.HTTPClient = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	s.Substitutions["YYYY-MM-DD"] = s.StartedAt.Format("2006-01-02")

	return s
}

func (state *state) log(tag string) {
	if state.LogTags {
		fmt.Printf("%#v\n\n", tag)
	}

	if state.LogState {
		fmt.Printf("%#v - %#v\n\n", tag, state)
	}
}
