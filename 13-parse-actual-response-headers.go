package main

import (
	"fmt"
	"sort"
)

func parseActualHeaders(context *context) bool {
	context.log("13 parseActualHeaders")

	var headerNames []string

	for headerName := range context.HTTPResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	sort.Strings(headerNames)

	for _, name := range headerNames {
		context.ActualResponseLineNumber++

		position :=
			fmt.Sprintf("actual-response:%v", context.ActualResponseLineNumber)

		headerText := "< " + name + ": " + context.HTTPResponse.Header.Get(name)

		headerLine := parse(position, headerText)

		context.ActualResponse.Lines =
			append(context.ActualResponse.Lines, headerLine)
	}

	return parseActualBlankLine(context)
}
