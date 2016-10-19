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
