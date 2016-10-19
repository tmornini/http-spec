package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func parseActualResponse(context *context) bool {
	context.log("10 parseActualResponse")

	version := context.HTTPResponse.Proto

	parts := strings.Split(context.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := parts[1]

	statusLine := fmt.Sprintf("< %s %s %s", version, statusCode, reasonPhrase)

	lines := []*line{parse(statusLine)}

	var headerNames []string

	actualResponse := context.HTTPResponse

	for headerName := range actualResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	sort.Strings(headerNames)

	for _, name := range headerNames {
		line := parse("< " + name + ": " + actualResponse.Header.Get(name))

		lines = append(lines, line)
	}

	line := parse("< ")

	lines = append(lines, line)

	scanner := bufio.NewScanner(context.HTTPResponse.Body)

	for scanner.Scan() {
		exitWithStatusOneIf(scanner.Err())

		line := parse("< " + scanner.Text())

		lines = append(lines, line)
	}

	context.HTTPResponse.Body.Close()

	context.ActualResponse = &response{
		Version:      version,
		StatusCode:   statusCode,
		ReasonPhrase: reasonPhrase,
		Lines:        lines,
	}

	return compareActualToExpected(context)
}
