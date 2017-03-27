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

	startedAt := time.Now()

	flag.StringVar(
		&prefix,
		"prefix",
		"",
		"prefix for request URLs",
	)

	flag.Parse()

	if prefix != "" {
		fmt.Fprintln(os.Stderr, "-prefix has been deprecated, please use absolute URIs in the request line")
		return
	} else {
		prefix = "http://localhost:80"
	}

	context := &context{
		LogFunctions:          false,
		LogContext:            false,
		URLPrefix:             prefix,
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
