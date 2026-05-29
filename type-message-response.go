package main

import "strings"

const dateHeaderPrefix = "< Date: "

func responseFromFile(context *context) (*response, error) {
	message, err := messageFromFile(context, "<")

	if err != nil {
		return nil, err
	}

	return &response{message}, nil
}

type response struct {
	*message
}

func (response *response) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, response.FirstLine.Content())

	for _, l := range response.HeaderLines {
		content := l.Content()

		if strings.HasPrefix(content, dateHeaderPrefix) {
			content =
				dateHeaderPrefix +
					regexpIdentifier +
					regexpIdentifier +
					":date" +
					regexpIdentifier
		}

		lineStrings = append(lineStrings, content)
	}

	lineStrings = append(lineStrings, response.BlankLine.Content())

	for _, l := range response.BodyLines {
		lineStrings = append(lineStrings, l.Content())
	}

	return strings.Join(lineStrings, "\n")
}
