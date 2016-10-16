package main

import (
	"regexp"
	"strings"
)

type line struct {
	IOPrefix   string
	Text       string
	RegexpName string
	Regexp     *regexp.Regexp
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

func (line *line) validate() {
	if line.isRequest() || line.isResponse() || line.isComment() {
		return
	}

	panic("lines must begin with <, > or #")
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

func (line *line) String() string {
	return line.Text
}
