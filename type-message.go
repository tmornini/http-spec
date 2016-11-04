package main

import (
	"io"
	"strings"
)

type message struct {
	FirstLine   *line
	HeaderLines []*line
	BlankLine   *line
	BodyLines   []*line
}

func messageFromFile(context *context) (*message, error) {
	firstLine, err := lineFromFile(context)

	if err != nil {
		return nil, err
	}

	var headerLine *line
	var headerLines []*line
	var emptyLine *line

	for {
		headerLine, err = lineFromFile(context)

		if err != nil {
			return nil, err
		}

		if headerLine.isEmpty() {
			emptyLine = headerLine

			break
		}

		headerLines = append(headerLines, headerLine)
	}

	var bodyLine *line
	var bodyLines []*line

	for {
		bodyLine, err = lineFromFile(context)

		if err == io.EOF || bodyLine.isBlank() {
			break
		}

		if err != nil {
			return nil, err
		}

		bodyLines = append(bodyLines, bodyLine)
	}

	return &message{
		firstLine,
		headerLines,
		emptyLine,
		bodyLines,
	}, nil
}

func (message *message) allHeaderAndBodyLines() []*line {
	allHeaderAndBodyLines := append([]*line{}, message.HeaderLines...)

	return append(allHeaderAndBodyLines, message.BodyLines...)
}

func (message *message) substitute(context *context) {
	for _, headerLine := range message.HeaderLines {
		headerLine.substitute(context)
	}

	for _, bodyLine := range message.BodyLines {
		bodyLine.substitute(context)
	}
}

func (message *message) Header() string {
	headerLineTexts := []string{}

	for _, headerLine := range message.HeaderLines {
		headerLineTexts = append(headerLineTexts, headerLine.Text)
	}

	return strings.Join(headerLineTexts, "\n")
}

func (message *message) Body() string {
	bodyLineTexts := []string{}

	for _, bodyLine := range message.BodyLines {
		bodyLineTexts = append(bodyLineTexts, bodyLine.Text)
	}

	return strings.Join(bodyLineTexts, "\n")
}
