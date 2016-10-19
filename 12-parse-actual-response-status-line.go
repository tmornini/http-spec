package main

import (
	"fmt"
	"strings"
)

func parseActualStatusLine(context *context) bool {
	context.log("12 parseActualResponseLine")

	version := context.HTTPResponse.Proto

	parts := strings.Split(context.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := parts[1]

	statusText := fmt.Sprintf("< %s %s %s", version, statusCode, reasonPhrase)

	context.ActualResponseLineNumber = 1

	position :=
		fmt.Sprintf("actual-response:%v", context.ActualResponseLineNumber)

	statusLine := parse(position, statusText)

	context.ActualResponse = &response{
		Version:      version,
		StatusCode:   statusCode,
		ReasonPhrase: reasonPhrase,
		Lines:        []*line{statusLine},
	}

	return parseActualHeaders(context)
}
