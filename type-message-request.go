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

	return &request{message}, err
}

type request struct {
	*message
}

func (request *request) Method() string {
	return strings.Split(request.FirstLine.Text, " ")[0]
}

func (request *request) Path() string {
	return strings.Split(request.FirstLine.Text, " ")[1]
}

func (request *request) Version() string {
	return strings.Split(request.FirstLine.Text, " ")[2]
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
