package main

import "strings"

func parseExpectedStatusLine(context *context) bool {
	context.log("06 parseExpectedStatusLine")

	context.Scanner.Scan()

	exitWithStatusOneIf(context.Scanner.Err())

	inputLine := context.Scanner.Text()

	statusLine := parse(inputLine)

	parts := strings.Split(statusLine.Text, " ")

	context.ExpectedResponse = &response{
		Version:      parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
		Lines:        []*line{statusLine},
	}

	return parseExpectedResponseLines(context)
}
