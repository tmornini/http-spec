package main

import "strings"

type response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
	Lines        []*line
}

func (response *response) compare(context *context, otherResponse *response) {
	for i, line := range response.Lines {
		line.compare(context, otherResponse.Lines[i])
	}
}

func (response *response) String() string {
	var lineStrings []string

	for _, line := range response.Lines {
		lineStrings = append(lineStrings, line.Text)
	}

	return strings.Join(lineStrings, "\n")
}
