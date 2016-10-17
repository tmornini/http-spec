package main

import "strings"

func parseExpectedStatusLine(context *context) bool {
	context.log("5 parseExpectedStatusLine")

	context.Scanner.Scan()

	if err := context.Scanner.Err(); err != nil {
		panic(err)
	}

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
