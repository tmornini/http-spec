package main

import (
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

func getMessageFromFile(state *state) (*message, error) {
	msg := &message{}
	var err error

	msg.FirstLine, err = getNextLineFromFile(state)

	if err != nil {
		return nil, err
	}

	for {
		headerLine, err := getNextLineFromFile(state)

		if err != nil {
			return nil, err
		}

		if headerLine.isEmpty() {
			msg.BlankLine = headerLine

			break
		}

		msg.HeaderLines = append(msg.HeaderLines, headerLine)
	}

	msg.Duration = time.Duration(0)

	for {
		bodyLine, err := getNextLineFromFile(state)

		if err == io.EOF || bodyLine.isBlank() {
			break
		}

		if err != nil {
			return nil, err
		}

		if bodyLine.isSleep() {
			msg.Duration, err = time.ParseDuration(bodyLine.Text)

			if err != nil {
				return nil, err
			}

			break
		}

		msg.BodyLines = append(msg.BodyLines, bodyLine)
	}

	return msg, nil
}

func (message *message) allLines() []*line {
	var allLines []*line

	allLines = append(allLines, message.FirstLine)
	allLines = append(allLines, message.HeaderLines...)
	allLines = append(allLines, message.BlankLine)
	allLines = append(allLines, message.BodyLines...)

	return allLines
}

func (message *message) substitute(state *state) {
	for _, line := range message.allLines() {
		line.substitute(state)
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
