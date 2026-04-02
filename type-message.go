package main

import (
	"fmt"
	"io"
	"strings"
	"time"
)

type message struct {
	FirstLine   *line
	HeaderLines []*line
	BlankLine   *line
	BodyLines   []*line
	Duration    time.Duration
}

func messageFromFile(
	context *context,
	expectedPrefix string,
) (*message, error) {
	firstLine, err := newLineFromFile(context)

	if err != nil {
		return nil, err
	}

	if firstLine.IOPrefix != expectedPrefix {
		return nil, fmt.Errorf(
			"expected %q prefix, got %q at line %d: %s",
			expectedPrefix,
			firstLine.IOPrefix,
			firstLine.LineNumber,
			firstLine.InputText,
		)
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

		if headerLine.IOPrefix != expectedPrefix {
			return nil, fmt.Errorf(
				"expected %q prefix, got %q at line %d: %s",
				expectedPrefix,
				headerLine.IOPrefix,
				headerLine.LineNumber,
				headerLine.InputText,
			)
		}

		headerLines = append(headerLines, headerLine)
	}

	var bodyLine *line
	var bodyLines []*line

	duration := time.Duration(0)

	for {
		bodyLine, err = newLineFromFile(context)

		if err == io.EOF || bodyLine.isBlank() {
			break
		}

		if err != nil {
			return nil, err
		}

		if bodyLine.isSleep() {
			duration, err = time.ParseDuration(bodyLine.Text)

			if err != nil {
				return nil, err
			}

			break
		}

		if bodyLine.IOPrefix != expectedPrefix {
			return nil, fmt.Errorf(
				"expected %q prefix, got %q at line %d: %s",
				expectedPrefix,
				bodyLine.IOPrefix,
				bodyLine.LineNumber,
				bodyLine.InputText,
			)
		}

		bodyLines = append(bodyLines, bodyLine)
	}

	return &message{
		firstLine,
		headerLines,
		emptyLine,
		bodyLines,
		duration,
	}, nil
}

func (message *message) allLines() []*line {
	var allLines []*line

	allLines = append(allLines, message.FirstLine)
	allLines = append(allLines, message.HeaderLines...)
	allLines = append(allLines, message.BlankLine)
	allLines = append(allLines, message.BodyLines...)

	return allLines
}

func (message *message) substitute(context *context) {
	for _, line := range message.allLines() {
		if errorHandler(context, line.substitute(context)) {
			return
		}
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
