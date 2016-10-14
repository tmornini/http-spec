package main

import "strings"

func parseStatusLine(context *context) {
	context.Scanner.Scan()

	statusLine := context.Scanner.Text()

	words := strings.Split(statusLine, " ")

	context.ExpectedResponse = &response{
		Version:      words[0],
		StatusCode:   words[1],
		ReasonPhrase: words[2],
	}

	parseResponseHeaders(context)
}
