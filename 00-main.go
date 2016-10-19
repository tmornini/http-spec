package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
)

const logFunctions = false
const logContext = false

const regexpIdentifier = "⧆"
const substitionIdentifier = "⧈"

func main() {
	if logFunctions {
		fmt.Println("00 main")
	}

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
		HTTPS:         https,
		HostName:      hostname,
		Pathnames:     flag.Args(),
		URIScheme:     uriScheme,
		Substitutions: map[string]string{},
		WaitGroup:     &sync.WaitGroup{},
	}

	processSpecFilesConcurrently(context)
}

func exitWithSucess() {
	os.Exit(0)
}

func exitWithFailure(message string) {
	os.Stderr.WriteString(message + "\n")
	os.Exit(1)
}

func exitWithFailureOn(err error) {
	if err != nil {
		exitWithFailure(err.Error())
	}
}
