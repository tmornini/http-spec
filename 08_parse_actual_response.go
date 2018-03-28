package main

import (
	"bufio"
	"errors"
	"fmt"
	"sort"
	"strings"
)

func parseActualResponse(state *state) error {
	state.log("parse-actual-response")

	message := &message{}
	var err error

	// Get Status Line
	message.FirstLine, err = getStatusLine(state)

	if err != nil {
		return err
	}

	// Get Header Lines
	message.HeaderLines, err = getHeaderLines(state)

	if err != nil {
		return err
	}

	// Get Blank Line
	state.ResponseLineNumber++
	message.BlankLine, err = newLineFromText(
		"response",
		state.ResponseLineNumber,
		"<",
	)

	if err != nil {
		return err
	}

	// Get body lines
	message.BodyLines, err = getBodyLines(state)

	if err != nil {
		return err
	}

	state.SpecTriplet.ActualResponse = &response{message}

	if state.SpecTriplet.isRequestOnly() {
		return fmt.Errorf("no expected response")
	}

	return compareResponse(state)
}

func getStatusLine(state *state) (*line, error) {
	state.ResponseLineNumber++

	version := state.HTTPResponse.Proto

	parts := strings.Split(state.HTTPResponse.Status, " ")

	if len(parts) < 2 {
		return nil, errors.New("invalid status line")
	}

	statusCode := parts[0]
	reasonPhrase := strings.Join(parts[1:], " ")

	statusLineText := "< " + version + " " + statusCode + " " + reasonPhrase

	statusLine, err := newLineFromText(
		"response",
		state.ResponseLineNumber,
		statusLineText,
	)

	if err != nil {
		return nil, err
	}

	return statusLine, nil
}

func getHeaderLines(state *state) ([]*line, error) {
	var err error

	var headerNames []string

	for headerName := range state.HTTPResponse.Header {
		headerNames = append(headerNames, headerName)
	}

	var headerLine *line
	var headerLines []*line

	sort.Strings(headerNames)

	for _, name := range headerNames {
		headerText := "< " + name + ": " + state.HTTPResponse.Header.Get(name)

		state.ResponseLineNumber++

		headerLine, err = newLineFromText(
			"response",
			state.ResponseLineNumber,
			headerText,
		)

		if err != nil {
			return nil, err
		}

		headerLines = append(headerLines, headerLine)
	}

	return headerLines, nil
}

func getBodyLines(state *state) ([]*line, error) {
	var err error

	scanner := bufio.NewScanner(state.HTTPResponse.Body)

	var bodyLine *line
	var bodyLines []*line

	for scanner.Scan() {
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}

		state.ResponseLineNumber++

		bodyLine, err = newLineFromText(
			"response",
			state.ResponseLineNumber,
			"< "+scanner.Text(),
		)

		if err != nil {
			return nil, err
		}

		bodyLines = append(bodyLines, bodyLine)
	}

	state.HTTPResponse.Body.Close()

	return bodyLines, nil
}
