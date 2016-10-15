package main

import "strings"

type response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
	Lines        []*line
}

func (response *response) String() string {
	assembly := response.Version + " " + response.StatusCode

	assembly += " " + response.ReasonPhrase + "\n"

	var inputLines []string

	for _, line := range response.Lines {
		inputLines = append(inputLines, line.Text)
	}

	assembly += strings.Join(inputLines, "\n")

	return assembly
}
