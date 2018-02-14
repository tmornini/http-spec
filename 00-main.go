package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

const regexpIdentifier = "⧆"
const substitutionIdentifier = "⧈"

func main() {
	startedAt := time.Now()

	var httpRetryDelay time.Duration
	var maxHTTPAttempts int
	var prefix string
	var skipTLSVerification bool

	defaultHTTPRetryDelay, err := time.ParseDuration("1s")
	defaultMaxHTTPAttempts := 180

	if err != nil {
		panic(err)
	}

	flag.DurationVar(
		&httpRetryDelay,
		"http-retry-delay",
		defaultHTTPRetryDelay,
		"delay between failed HTTP requests",
	)

	flag.IntVar(
		&maxHTTPAttempts,
		"max-http-attempts",
		defaultMaxHTTPAttempts,
		"maximum number of attempts per HTTP request",
	)

	flag.StringVar(
		&prefix,
		"prefix",
		"",
		"prefix for request URLs",
	)

	flag.BoolVar(
		&skipTLSVerification,
		"skip-tls-verification",
		false,
		"skip TLS verification (hostname mismatch, self-signed certifications, etc.)",
	)

	flag.Parse()

	if prefix != "" {
		_, err := fmt.Fprintln(os.Stderr, "-prefix has been deprecated, please use absolute URIs in the request line")

		if err != nil {
			panic(err)
		}

		os.Exit(1)
	} else {
		prefix = "http://localhost:80"
	}

	context := &context{
		HTTPRetryDelay:        httpRetryDelay,
		LogContext:            false,
		LogFunctions:          false,
		MaxHTTPAttempts:       maxHTTPAttempts,
		Pathnames:             flag.Args(),
		ResultGathererChannel: make(chan context),
		SkipTLSVerification:   skipTLSVerification,
		StartedAt:             startedAt,
		URLPrefix:             prefix,
		WaitGroup:             &sync.WaitGroup{},
	}

	context.log("00 main")

	go resultGatherer(*context)

	specFileScatter(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
