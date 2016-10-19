package main

import "strings"

type request struct {
	Verb    string
	Path    string
	Version string
	Lines   []*line
}

func (request *request) String() string {
	var lineStrings []string

	for _, line := range request.Lines {
		lineStrings = append(lineStrings, line.Text)
	}

	return strings.Join(lineStrings, "\n")
}
