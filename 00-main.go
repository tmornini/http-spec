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

	var hostname string
	var httpRetryDelay time.Duration
	var maxHTTPAttempts int
	var scheme string
	var skipTLSVerification bool

	defaultHTTPRetryDelay, err := time.ParseDuration("1s")

	if err != nil {
		panic(err)
	}

	flag.StringVar(
		&hostname,
		"hostname",
		"",
		"hostname",
	)

	flag.DurationVar(
		&httpRetryDelay,
		"http-retry-delay",
		defaultHTTPRetryDelay,
		"delay between failed HTTP requests",
	)

	flag.IntVar(
		&maxHTTPAttempts,
		"max-http-attempts",
		180,
		"maximum number of attempts per HTTP request",
	)

	flag.StringVar(
		&scheme,
		"scheme",
		"",
		"scheme (http/https)",
	)

	flag.BoolVar(
		&skipTLSVerification,
		"skip-tls-verification",
		false,
		"skip TLS verification (hostname mismatch, self-signed certifications, etc.)",
	)

	flag.Parse()

	context := &context{
		Hostname:              hostname,
		HTTPRetryDelay:        httpRetryDelay,
		LogContext:            false,
		LogFunctions:          false,
		MaxHTTPAttempts:       maxHTTPAttempts,
		Pathnames:             flag.Args(),
		ResultGathererChannel: make(chan context),
		SkipTLSVerification:   skipTLSVerification,
		StartedAt:             startedAt,
		Scheme:                scheme,
		WaitGroup:             &sync.WaitGroup{},
	}

	context.log("00 main")

	go resultGatherer(*context)

	specFileScatter(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
