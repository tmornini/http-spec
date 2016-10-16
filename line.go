package main

import (
	"fmt"
	"regexp"
	"strings"
)

type line struct {
	IOPrefix   string
	Text       string
	RegexpName string
	Regexp     *regexp.Regexp
}

func (line *line) isRequest() bool {
	return strings.HasPrefix(string(line.IOPrefix[0]), ">")
}

func (line *line) isResponse() bool {
	return strings.HasPrefix(string(line.IOPrefix[0]), "<")
}

func (line *line) isComment() bool {
	return strings.HasPrefix(string(line.IOPrefix[0]), "#")
}

func parse(inputLine string) *line {
	ioPrefix, text := separate(inputLine)

	line := &line{
		IOPrefix: ioPrefix,
		Text:     text,
	}

	line.validate()

	return line
}

func separate(inputLine string) (string, string) {
	length := len(inputLine)

	switch length {
	case 0:
		return "", ""
	case 1:
		return string(inputLine[0]), ""
	case 2:
		return string(inputLine[0:1]), ""
	default:
		return string(inputLine[0:1]), string(inputLine[2:length])
	}
}

func (line *line) substitute(context *context) {
	parts := strings.Split(line.Text, substitionIdentifier)

	switch len(parts) {
	case 1:
		return
	case 3:
		line.Text = parts[0] + context.Substitutions[parts[1]] + parts[2]
		fmt.Printf("substituted -> %#v\n", line.Text)
	default:
		panic(
			fmt.Sprintf(
				"substition must be formed %scapture-name%s",
				substitionIdentifier,
				substitionIdentifier,
			),
		)
	}
}

func (line *line) validate() {
	if line.isRequest() || line.isResponse() || line.isComment() {
		return
	}

	panic("lines must begin with <, > or #")
}

func (line *line) String() string {
	return line.Text
}
