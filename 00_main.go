package main

import (
	"flag"
	"os"
	"time"
)

func main() {
	var httpRetryDelay time.Duration
	flag.DurationVar(
		&httpRetryDelay,
		"http-retry-delay",
		time.Second,
		"delay between failed HTTP requests",
	)

	var maxHTTPAttempts int
	flag.IntVar(
		&maxHTTPAttempts,
		"max-http-attempts",
		10,
		"maximum number of attempts per HTTP request",
	)

	var skipTLSVerification bool
	flag.BoolVar(
		&skipTLSVerification,
		"skip-tls-verification",
		false,
		"skip TLS verification (hostname mismatch, self-signed certifications, etc.)",
	)

	flag.Parse()

	state := NewState(skipTLSVerification)
	state.HTTPRetryDelay = httpRetryDelay
	state.MaxHTTPAttempts = maxHTTPAttempts
	state.HTSFFilePaths = flag.Args()

	// state.LogTags = true
	// state.LogState = true

	// Setup the pipeline
	p1 := iterateSpecFiles(*state)

	p2 := fanIn(
		processSpecFile(p1),
		processSpecFile(p1),
		processSpecFile(p1),
		processSpecFile(p1),
		processSpecFile(p1),
	)

	p3 := fanIn(
		iterateSpecTriplet(p2),
		iterateSpecTriplet(p2),
		iterateSpecTriplet(p2),
		iterateSpecTriplet(p2),
		iterateSpecTriplet(p2),
	)

	p4 := fanIn(
		parseResults(p3),
		parseResults(p3),
		parseResults(p3),
		parseResults(p3),
		parseResults(p3),
	)

	p5 := outputResults(p4)

	os.Exit(<-p5)
}
