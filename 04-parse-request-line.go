package main

import "strings"

func parseRequestLine(context *context) bool {
	context.log("04 parseRequestLine")

	requestLine := context.File.readLine()

	if len(requestLine.InputLine) == 0 {
		return true
	}

	parts := strings.Split(requestLine.Text, " ")

	context.Request = &request{
		Method:  parts[0],
		Path:    parts[1],
		Version: parts[2],
		Lines:   []*line{requestLine},
	}

	return parseRequestLines(context)
}
