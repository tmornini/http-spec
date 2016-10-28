package main

import (
	"flag"
	"sync"
)

const regexpIdentifier = "⧆"
const substitionIdentifier = "⧈"

func main() {
	var https bool
	var hostname string

	flag.BoolVar(&https, "https", false, "use HTTPS")
	flag.StringVar(&hostname, "hostname", "localhost", "`hostname` to test")

	flag.Parse()

	var uriScheme string

	if https {
		uriScheme = "https://"
	} else {
		uriScheme = "http://"
	}

	context := &context{
		LogFunctions:          false,
		LogContext:            false,
		HTTPS:                 https,
		HostName:              hostname,
		Pathnames:             flag.Args(),
		URIScheme:             uriScheme,
		WaitGroup:             &sync.WaitGroup{},
		ResultGathererChannel: make(chan context),
	}

	context.log("00 main")

	// not in WaitGroup, terminated at close of ResultGathererChannel below
	go resultGatherer(*context)

	specFileScatter(context)

	context.WaitGroup.Wait()

	close(context.ResultGathererChannel)

	select {}
}
