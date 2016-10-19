package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"
)

func translateResponse(context *context) error {
	context.log("08 translate-response")

	message := &message{}

	context.ResponseLineNumber = 1

	version, statusCode, reasonPhrase, statusLine, err :=
		translateResponseStatusLine(context)

	if err != nil {
		return err
	}

	message.FirstLine = statusLine

	headerLines, err := translateResponseHeader(context)

	if err != nil {
		return err
	}

	message.HeaderLines = headerLines

	message.BlankLine, err =
		lineFromText("response", context.ResponseLineNumber, "<")

	if err != nil {
		return err
	}

	bodyLines, err := translateResponseBody(context)

	if err != nil {
		return err
	}

	message.BodyLines = bodyLines

	context.SpecTriplet.ActualResponse = &response{
		message,
		version,
		statusCode,
		reasonPhrase,
	}

	if context.SpecTriplet.ExpectedResponse == nil {
		fmt.Println(context.SpecTriplet.ActualResponse)

		return fmt.Errorf("no expected response")
	}

	return compareResponses(context)
}

func translateResponseStatusLine(context *context) (string, string, string, *line, error) {
	version := context.HTTPResponse.Proto

	parts := strings.Split(context.HTTPResponse.Status, " ")

	statusCode := parts[0]
	reasonPhrase := parts[1]

	statusLineText := "< " + version + " " + statusCode + " " + reasonPhrase

	statusLine, err :=
		lineFromText("response", context.ResponseLineNumber, statusLineText)

	if err != nil {
		return "", "", "", nil, err
	}

	context.ResponseLineNumber++

	return version, statusCode, reasonPhrase, statusLine, nil
}

func translateResponseHeader(context *context) ([]*line, error) {
	var headerNames []string

	for headerName := range context.HTTPResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	sort.Strings(headerNames)

	var headerLines []*line

	for _, name := range headerNames {
		headerText := "< " + name + ": " + context.HTTPResponse.Header.Get(name)

		headerLine, err :=
			lineFromText("response", context.ResponseLineNumber, headerText)

		if err != nil {
			return nil, err
		}

		headerLines =
			append(headerLines, headerLine)

		context.ResponseLineNumber++
	}

	return headerLines, nil
}

func translateResponseBody(context *context) ([]*line, error) {
	scanner := bufio.NewScanner(context.HTTPResponse.Body)

	var bodyLines []*line

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}

		bodyLine, err :=
			lineFromText("response", context.ResponseLineNumber, "< "+scanner.Text())

		if err != nil {
			return nil, err
		}

		bodyLines =
			append(bodyLines, bodyLine)

		context.ResponseLineNumber++
	}

	context.HTTPResponse.Body.Close()

	return bodyLines, nil
}
