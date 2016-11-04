package main

import (
	"fmt"
	"strings"
)

func responseFromFile(context *context) (*response, error) {
	message, err := messageFromFile(context)

	if err != nil {
		return nil, err
	}

	statusLine := message.FirstLine

	parts := strings.Split(statusLine.Text, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("malformed status-line: %s", statusLine.String())
	}

	return &response{message}, nil
}

type response struct {
	*message
}

func (response *response) Version() string {
	return strings.Split(response.FirstLine.Text, " ")[0]
}

func (response *response) StatusCode() string {
	return strings.Split(response.FirstLine.Text, " ")[1]
}

func (response *response) ReasonPhrase() string {
	return strings.Split(response.FirstLine.Text, " ")[2]
}

func (response *response) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, response.FirstLine.Content())

	for _, l := range response.HeaderLines {
		content := l.Content()

		if content[0:8] == "< Date: " {
			content =
				content[0:8] +
					regexpIdentifier +
					regexpIdentifier +
					":date" +
					regexpIdentifier
		}

		lineStrings = append(lineStrings, content)
	}

	lineStrings = append(lineStrings, response.BlankLine.Content())

	if response.BodyLines != nil {
		for _, l := range response.BodyLines {
			lineStrings = append(lineStrings, l.Content())
		}
	}

	return strings.Join(lineStrings, "\n")
}
