package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func actualResponseReceiver(context *context) {
	context.log("08 actual-response-receiver")

	message := &message{}

	responseLineNumber := 1

	version := context.HTTPResponse.Proto

	parts := strings.Split(context.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := parts[1]

	statusLineText := "< " + version + " " + statusCode + " " + reasonPhrase

	statusLine, err :=
		newLineFromText("response", responseLineNumber, statusLineText)

	if errorHandler(context, err) {
		return
	}

	responseLineNumber++

	message.FirstLine = statusLine

	var headerNames []string

	for headerName := range context.HTTPResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	var headerLine *line
	var headerLines []*line

	sort.Strings(headerNames)

	for _, name := range headerNames {
		headerText := "< " + name + ": " + context.HTTPResponse.Header.Get(name)

		headerLine, err =
			newLineFromText("response", responseLineNumber, headerText)

		if errorHandler(context, err) {
			return
		}

		headerLines =
			append(headerLines, headerLine)

		responseLineNumber++
	}

	message.HeaderLines = headerLines

	message.BlankLine, err =
		newLineFromText("response", responseLineNumber, "<")

	if errorHandler(context, err) {
		return
	}

	responseLineNumber++

	scanner := bufio.NewScanner(context.HTTPResponse.Body)

	var bodyLine *line
	var bodyLines []*line

	for scanner.Scan() {
		if errorHandler(context, scanner.Err()) {
			return
		}

		bodyLine, err =
			newLineFromText("response", responseLineNumber, "< "+scanner.Text())

		if errorHandler(context, err) {
			return
		}

		bodyLines =
			append(bodyLines, bodyLine)

		responseLineNumber++
	}

	context.HTTPResponse.Body.Close()

	message.BodyLines = bodyLines

	context.SpecTriplet.ActualResponse = &response{message}

	if context.SpecTriplet.isRequestOnly() {
		context.Err = fmt.Errorf("no expected response")

		return
	}

	responseComparitor(context)
}
