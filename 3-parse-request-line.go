package main

import "strings"

func parseRequestLine(context *context) {
	context.Scanner.Scan()

	requestLine := context.Scanner.Text()

	parts := strings.Split(requestLine, " ")

	context.Request = &request{
		Verb:    parts[0],
		Path:    parts[1],
		Version: parts[2],
	}

	parseRequestHeaders(context)
}
