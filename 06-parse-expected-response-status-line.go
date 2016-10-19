package main

import "strings"

func parseExpectedStatusLine(context *context) bool {
	context.log("06 parseExpectedStatusLine")

	statusLine := context.File.readLine()

	parts := strings.Split(statusLine.Text, " ")

	context.ExpectedResponse = &response{
		Version:      parts[0],
		StatusCode:   parts[1],
		ReasonPhrase: parts[2],
		Lines:        []*line{statusLine},
	}

	return parseExpectedResponseLines(context)
}
