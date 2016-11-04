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

	requestLine := message.FirstLine

	parts := strings.Split(requestLine.Text, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("%s: malformed request-line", requestLine.String())
	}

	return &response{
		message,
		parts[0],
		parts[1],
		parts[2],
	}, nil
}

type response struct {
	*message
	Version      string
	StatusCode   string
	ReasonPhrase string
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
