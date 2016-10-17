package main

import "strings"

func parseExpectedStatusLine(context *context) bool {
	context.log("05 parseExpectedStatusLine")

	context.Scanner.Scan()

	panicOn(context.Scanner.Err())

	inputLine := context.Scanner.Text()

	statusLine := parse(inputLine)

	parts := strings.Split(statusLine.Text, " ")

	context.ExpectedResponse = &response{
		Version:      parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
	}

	return parseExpectedResponseHeaders(context)
}
