package main

import (
	"flag"
	"strconv"
	"strings"
	"sync"
	"time"
)

const regexpIdentifier = "⧆"
const substitutionIdentifier = "⧈"
const defaultMaxHTTPAttempts = 180

func main() {
	startedAt := time.Now()

	var hostname string
	var httpRetryDelay time.Duration
	var matchersPath string
	var maxHTTPAttempts int
	var retryStatusCodesStr string
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

	flag.StringVar(
		&matchersPath,
		"matchers",
		"",
		"path to custom matchers JSON file",
	)

	flag.IntVar(
		&maxHTTPAttempts,
		"max-http-attempts",
		defaultMaxHTTPAttempts,
		"maximum number of attempts per HTTP request",
	)

	flag.StringVar(
		&retryStatusCodesStr,
		"retry-status-codes",
		"",
		"comma-separated HTTP status codes to retry on (e.g. 502,503)",
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
		"skip TLS certificate verification",
	)

	flag.Parse()

	var retryStatusCodes []int

	if retryStatusCodesStr != "" {
		for _, s := range strings.Split(retryStatusCodesStr, ",") {
			code, err := strconv.Atoi(strings.TrimSpace(s))

			if err != nil {
				panic("invalid status code: " + s)
			}

			retryStatusCodes = append(retryStatusCodes, code)
		}
	}

	context := &context{
		Hostname:              hostname,
		HTTPRetryDelay:        httpRetryDelay,
		LogContext:            false,
		LogFunctions:          false,
		Matchers:              make(map[string]string),
		MatchersPath:          matchersPath,
		MaxHTTPAttempts:       maxHTTPAttempts,
		Pathnames:             flag.Args(),
		ResultGathererChannel: make(chan context),
		RetryStatusCodes:      retryStatusCodes,
		SkipTLSVerification:   skipTLSVerification,
		StartedAt:             startedAt,
		Scheme:                scheme,
		WaitGroup:             &sync.WaitGroup{},
	}

	err = loadMatchers(context)

	if err != nil {
		panic(err)
	}

	context.log("00 main")

	go resultGatherer(*context)

	specFileScatterer(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
