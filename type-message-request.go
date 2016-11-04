package main

import (
	"fmt"
	"strings"
)

func requestFromFile(context *context) (*request, error) {
	message, err := messageFromFile(context)

	if err != nil {
		return nil, err
	}

	requestLine := message.FirstLine

	parts := strings.Split(requestLine.Text, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("%s: malformed request-line", requestLine.String())
	}

	return &request{
		message,
		parts[0],
		parts[1],
		parts[2],
	}, err
}

type request struct {
	*message
	Method  string
	Path    string
	Version string
}

func (request *request) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, request.FirstLine.String())

	for _, l := range request.HeaderLines {
		lineStrings = append(lineStrings, l.String())
	}

	lineStrings = append(lineStrings, request.BlankLine.String())

	if request.BodyLines != nil {
		for _, l := range request.BodyLines {
			lineStrings = append(lineStrings, l.String())
		}
	}

	return strings.Join(lineStrings, "\n")
}
