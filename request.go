package main

import "strings"

type request struct {
	Verb    string
	Path    string
	Version string
	Lines   []*line
}

func (request *request) String() string {
	assembly := request.Verb + " " + request.Path + " " + request.Version + "\n"

	var lineStrings []string

	for _, line := range request.Lines {
		lineStrings = append(lineStrings, line.String())
	}

	assembly += strings.Join(lineStrings, "\n")

	return assembly
}
