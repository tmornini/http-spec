package main

import (
	"flag"
	"fmt"
)

func processCLIArguments() {
	if logFunctions {
		fmt.Println("00 processCLIArguments")
	}

	var https bool
	var hostname string

	flag.BoolVar(&https, "https", false, "use HTTPS")

	flag.StringVar(&hostname, "hostname", "localhost", "`hostname` to test")

	flag.Parse()

	processFiles(https, hostname, flag.Args())
}
