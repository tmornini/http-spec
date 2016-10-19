package main

import "sort"

func parseActualHeaders(context *context) bool {
	context.log("12 parseActualHeaders")

	var headerNames []string

	for headerName := range context.HTTPResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	sort.Strings(headerNames)

	for _, name := range headerNames {
		line := parse("< " + name + ": " + context.HTTPResponse.Header.Get(name))

		context.ActualResponse.Lines = append(context.ActualResponse.Lines, line)
	}

	return parseActualBlankLine(context)
}
