package main

import (
	"bufio"
	"flag"
)

type context struct {
	Hostname string
	HTTPS    bool
	Pathname string
	Scanner  *bufio.Scanner
	Verb     string
	Path     string
	Version  string
	Status   string
}

func main() {
	var hostname string
	var https bool

	flag.StringVar(&hostname, "hostname", "localhost", "`hostname` to test")

	flag.BoolVar(&https, "https", false, "use HTTPS")

	flag.Parse()

	for _, pathname := range flag.Args() {
		parseFile(&context{
			Hostname: hostname,
			HTTPS:    https,
			Pathname: pathname,
		})
	}
}
