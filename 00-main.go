package main

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const regexpIdentifier = "⧆"
const substitutionIdentifier = "⧈"
const defaultMaxHTTPAttempts = 180
const defaultHTTPRetryDelay = 1 * time.Second
const defaultHTTPTimeout = 30 * time.Second

func main() {
	startedAt := time.Now()

	var hostname string
	var httpRetryDelay time.Duration
	var httpTimeout time.Duration
	var matchersPath string
	var maxHTTPAttempts int
	var retryStatusCodesStr string
	var scheme string
	var showSubstitutions bool
	var skipTLSVerification bool

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

	flag.DurationVar(
		&httpTimeout,
		"http-timeout",
		defaultHTTPTimeout,
		"timeout per HTTP request",
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
		&showSubstitutions,
		"show-substitutions",
		false,
		"show substitutions table on success (always shown on failure)",
	)

	flag.BoolVar(
		&skipTLSVerification,
		"skip-tls-verification",
		false,
		"skip TLS certificate verification",
	)

	flag.Parse()

	if len(flag.Args()) == 0 {
		os.Exit(0)
	}

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
		HTTPTimeout:           httpTimeout,
		LogContext:            false,
		LogFunctions:          false,
		Matchers:              make(map[string]string),
		MatchersPath:          matchersPath,
		MaxHTTPAttempts:       maxHTTPAttempts,
		Pathnames:             flag.Args(),
		ResultGathererChannel: make(chan context),
		RetryStatusCodes:      retryStatusCodes,
		ShowSubstitutions:     showSubstitutions,
		SkipTLSVerification:   skipTLSVerification,
		StartedAt:             startedAt,
		Scheme:                scheme,
		WaitGroup:             &sync.WaitGroup{},
	}

	err := loadMatchers(context)

	if err != nil {
		panic(err)
	}

	context.enterStage("00 main")

	go resultGatherer(*context)

	specFileScatterer(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
