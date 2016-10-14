package main

import "strings"

func parseExpectedResponseStatusLine(context *context) {
	context.Scanner.Scan()

	statusLine := context.Scanner.Text()

	parts := strings.Split(statusLine, " ")

	context.ExpectedResponse = &response{
		Version:      parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
	}

	parseExpectedResponseHeaders(context)
}
