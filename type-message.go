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
	firstLine, err := newLineFromFile(context)

	if err != nil {
		return nil, err
	}

	var headerLine *line
	var headerLines []*line
	var emptyLine *line

	for {
		headerLine, err = newLineFromFile(context)

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
		bodyLine, err = newLineFromFile(context)

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

func (message *message) allLines() []*line {
	var allLines []*line

	allLines = append(allLines, message.FirstLine)
	allLines = append(allLines, message.HeaderLines...)
	allLines = append(allLines, message.BodyLines...)

	return allLines
}

func (message *message) allHeaderAndBodyLines() []*line {
	var allHeaderAndBodyLines []*line

	allHeaderAndBodyLines = append(allHeaderAndBodyLines, message.HeaderLines...)
	allHeaderAndBodyLines = append(allHeaderAndBodyLines, message.BodyLines...)

	return allHeaderAndBodyLines
}

func (message *message) substitute(context *context) {
	for _, line := range message.allLines() {
		line.substitute(context)
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
