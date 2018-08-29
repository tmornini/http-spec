package main

import (
	"flag"
	"sync"
	"time"
)

const regexpIdentifier = "⧆"
const substitutionIdentifier = "⧈"

func main() {
	startedAt := time.Now()

	var httpRetryDelay time.Duration
	var maxHTTPAttempts int
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

	flag.BoolVar(
		&skipTLSVerification,
		"skip-tls-verification",
		false,
		"skip TLS verification (hostname mismatch, self-signed certifications, etc.)",
	)

	flag.Parse()

	context := &context{
		HTTPRetryDelay:        httpRetryDelay,
		LogContext:            false,
		LogFunctions:          false,
		MaxHTTPAttempts:       maxHTTPAttempts,
		Pathnames:             flag.Args(),
		ResultGathererChannel: make(chan context),
		SkipTLSVerification:   skipTLSVerification,
		StartedAt:             startedAt,
		WaitGroup:             &sync.WaitGroup{},
	}

	context.log("00 main")

	go resultGatherer(*context)

	specFileScatter(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
