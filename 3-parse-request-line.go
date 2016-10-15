package main

import "strings"

func parseRequestLine(context *context) {
	context.Scanner.Scan()

	inputLine := context.Scanner.Text()

	requestLine := parse(inputLine)

	parts := strings.Split(requestLine.Text, " ")

	context.Request = &request{
		Verb:    parts[0],
		Path:    parts[1],
		Version: parts[2],
	}

	parseRequestLines(context)
}
