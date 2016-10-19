package main

import (
	"fmt"
	"os"
	"sync"
)

func processFiles(https bool, hostname string, pathnames []string) {
	if logFunctions {
		fmt.Println("02 processFiles")
	}

	if logContext {
		fmt.Printf("https: %v, hostname: %v", https, hostname)
	}

	var uriScheme string

	if https {
		uriScheme = "https://"
	} else {
		uriScheme = "http://"
	}

	waitGroup := &sync.WaitGroup{}

	for _, pathname := range pathnames {
		waitGroup.Add(1)

		context := context{
			HTTPS:     https,
			HostName:  hostname,
			URIScheme: uriScheme,
			Pathname:  pathname,
			WaitGroup: waitGroup,
		}

		go processFile(context)
	}

	waitGroup.Wait()

	os.Exit(0)
}
