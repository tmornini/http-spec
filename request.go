package main

import "strings"

type request struct {
	Prefix      string
	Verb        string
	Path        string
	Version     string
	HeaderLines *[]string
	BodyLines   *[]string
}

func (request *request) String() string {
	assembly := request.Verb + " " + request.Path + " " + request.Version + "\n"

	assembly += strings.Join(*request.HeaderLines, "\n")

	assembly += "\n"

	assembly += strings.Join(*request.BodyLines, "\n")

	return assembly
}
