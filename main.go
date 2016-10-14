package main

import "flag"

func main() {
	var https bool
	var hostname string

	flag.BoolVar(&https, "https", false, "use HTTPS")

	flag.StringVar(&hostname, "hostname", "localhost", "`hostname` to test")

	flag.Parse()

	processFiles(https, hostname, flag.Args())
}
