package main

import (
	"fmt"
	"net/http"
	"strings"
)

func requestFromFile(context *context) (*request, error) {
	message, err := messageFromFile(context, ">")

	if err != nil {
		return nil, err
	}

	request := &request{
		message:  message,
		Hostname: context.Hostname,
		Scheme:   context.Scheme,
	}

	err = request.validate()

	if err != nil {
		return nil, err
	}

	return request, nil
}

func (request *request) validate() error {
	if len(strings.Split(request.FirstLine.Text, " ")) < 2 {
		return fmt.Errorf(
			"%s request line must be \"METHOD URI\": %q",
			request.FirstLine.Location(),
			request.FirstLine.Text,
		)
	}

	for _, headerLine := range request.HeaderLines {
		if !strings.Contains(headerLine.Text, ":") {
			return fmt.Errorf(
				"%s request header must contain \":\": %q",
				headerLine.Location(),
				headerLine.Text,
			)
		}
	}

	return nil
}

type request struct {
	*message
	Hostname string
	Scheme   string
}

func (request *request) Method() string {
	return strings.Split(request.FirstLine.Text, " ")[0]
}

func (request *request) uri() string {
	return strings.Split(request.FirstLine.Text, " ")[1]
}

func (request *request) HTTPHeaders() http.Header {
	headers := http.Header{}

	for _, headerLine := range request.HeaderLines {
		parts := strings.SplitN(headerLine.Text, ":", 2)

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		headers.Add(key, value)
	}

	return headers
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
		return request.uri()
	}

	return request.Scheme + "://" + hostname + request.uri()
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
