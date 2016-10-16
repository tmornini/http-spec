package main

import "strings"

func parseExpectedStatusLine(context *context) {
	context.Scanner.Scan()

	inputLine := context.Scanner.Text()

	statusLine := parse(inputLine)

	parts := strings.Split(statusLine.Text, " ")

	context.ExpectedResponse = &response{
		Version:      parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
	}

	parseExpectedResponseHeaders(context)
}
