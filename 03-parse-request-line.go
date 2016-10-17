package main

import "strings"

func parseRequestLine(context *context) bool {
	context.log("3 parseRequestLine")

	context.Scanner.Scan()

	inputLine := context.Scanner.Text()

	panicOn(context.Scanner.Err())

	if len(inputLine) == 0 {
		return true
	}

	requestLine := parse(inputLine)

	parts := strings.Split(requestLine.Text, " ")

	context.Request = &request{
		Verb:    parts[0],
		Path:    parts[1],
		Version: parts[2],
	}

	return parseRequestLines(context)
}
