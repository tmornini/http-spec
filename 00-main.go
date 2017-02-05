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
	var prefix string
	var skipTLSVerification bool

	startedAt := time.Now()

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
		LogFunctions:          false,
		LogContext:            false,
		URLPrefix:             prefix,
		SkipTLSVerification:   skipTLSVerification,
		Pathnames:             flag.Args(),
		WaitGroup:             &sync.WaitGroup{},
		ResultGathererChannel: make(chan context),
		StartedAt:             startedAt,
	}

	context.log("00 main")

	go resultGatherer(*context)

	specFileScatter(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
