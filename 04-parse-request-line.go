package main

import "strings"

func parseRequestLine(context *context) bool {
	context.log("04 parseRequestLine")

	context.Scanner.Scan()

	inputLine := context.Scanner.Text()

	exitWithStatusOneIf(context.Scanner.Err())

	if len(inputLine) == 0 {
		return true
	}

	requestLine := parse(inputLine)

	parts := strings.Split(requestLine.Text, " ")

	context.Request = &request{
		Verb:    parts[0],
		Path:    parts[1],
		Version: parts[2],
		Lines:   []*line{requestLine},
	}

	return parseRequestLines(context)
}
