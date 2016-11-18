package main

import (
	"flag"
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
		"http://localhost:80",
		"prefix for request URLs",
	)

	flag.Parse()

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
