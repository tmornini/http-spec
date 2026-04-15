package main

import "strings"

func requestFromFile(context *context) (*request, error) {
	message, err := messageFromFile(context, ">")

	if err != nil {
		return nil, err
	}

	parts := strings.Split(message.FirstLine.Text, " ")

	return &request{
		message:  message,
		Hostname: context.Hostname,
		Scheme:   context.Scheme,
		method:   parts[0],
		uri:      parts[1],
	}, nil
}

type request struct {
	*message
	Hostname string
	Scheme   string
	method   string
	uri      string
}

func (request *request) Method() string {
	return request.method
}

func (request *request) hostHeader() string {
	for _, headerLine := range request.HeaderLines {
		parts := strings.SplitN(headerLine.Text, ":", 2)

		if strings.TrimSpace(parts[0]) == "Host" {
			return strings.TrimSpace(parts[1])
		}
	}

	return ""
}

func (request *request) URL() string {
	hostname := request.Hostname

	if hostname == "" {
		hostname = request.hostHeader()
	}

	if hostname == "" {
		return request.uri
	}

	return request.Scheme + "://" + hostname + request.uri
}

func (request *request) String() string {
	lineStrings := []string{}

	lineStrings = append(lineStrings, request.FirstLine.String())

	for _, l := range request.HeaderLines {
		lineStrings = append(lineStrings, l.String())
	}

	lineStrings = append(lineStrings, request.BlankLine.String())

	for _, l := range request.BodyLines {
		lineStrings = append(lineStrings, l.String())
	}

	return strings.Join(lineStrings, "\n")
}
