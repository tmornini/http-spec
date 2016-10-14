package main

import "strings"

type response struct {
	Version      string
	StatusCode   string
	ReasonPhrase string
	HeaderLines  *[]*headerLine
	BodyLines    *[]string
}

func (response *response) String() string {
	assembly := response.Version + " " + response.StatusCode

	assembly += " " + response.ReasonPhrase + "\n"

	var headerLineInputs []string

	for _, headerLine := range *response.HeaderLines {
		headerLineInputs = append(headerLineInputs, headerLine.Input)
	}

	assembly += strings.Join(headerLineInputs, "\n")

	assembly += "\n"

	// assembly += strings.Join(*response.BodyLines, "\n")

	return assembly
}
