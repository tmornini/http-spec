package main

import "strings"

func parseRequestLine(context *context) {
	context.Scanner.Scan()

	requestLine := context.Scanner.Text()

	words := strings.Split(requestLine, " ")

	context.Request = &request{
		Verb:    words[0],
		Path:    words[1],
		Version: words[2],
	}

	parseRequestHeaders(context)
}
