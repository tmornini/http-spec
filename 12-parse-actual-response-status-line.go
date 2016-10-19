package main

import (
	"fmt"
	"strings"
)

func parseActualStatusLine(context *context) bool {
	context.log("11 parseActualResponseLine")

	version := context.HTTPResponse.Proto

	parts := strings.Split(context.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := parts[1]

	statusLine := fmt.Sprintf("< %s %s %s", version, statusCode, reasonPhrase)

	context.ActualResponse = &response{
		Version:      version,
		StatusCode:   statusCode,
		ReasonPhrase: reasonPhrase,
		Lines:        []*line{parse(statusLine)},
	}

	return parseActualHeaders(context)
}
