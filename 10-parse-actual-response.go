package main

import (
	"bufio"
	"sort"
	"strings"
)

func parseActualResponse(context *context) bool {
	context.log("10 parseRequest")

	version := context.HTTPResponse.Proto

	parts := strings.Split(context.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := parts[1]

	var lines []*line

	var names []string

	actualResponse := context.HTTPResponse

	for name := range actualResponse.Header {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		line := parse("< " + name + ": " + actualResponse.Header.Get(name))

		lines = append(lines, line)
	}

	scanner := bufio.NewScanner(context.HTTPResponse.Body)

	for scanner.Scan() {
		panicOn(scanner.Err())

		line := parse("< " + context.Scanner.Text())

		lines = append(lines, line)
	}

	context.HTTPResponse.Body.Close()

	context.ActualResponse = &response{
		Version:      version,
		StatusCode:   statusCode,
		ReasonPhrase: reasonPhrase,
		Lines:        lines,
	}

	return false
}
