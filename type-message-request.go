package main

import (
	"fmt"
	"strings"
)

func requestFromFile(context *context) (*request, error) {
	message, err := messageFromFile(context)

	if err != nil {
		return nil, err
	}

	statusLine := message.FirstLine

	parts := strings.Split(statusLine.Text, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("malformed status-line: %s", statusLine.String())
	}

	return &request{
		message,
		parts[0],
		parts[1],
		parts[2],
	}, err
}

type request struct {
	*message
	Method  string
	Path    string
	Version string
}
