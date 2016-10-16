package main

import (
	"regexp"
	"strings"
)

type line struct {
	Prefix     string
	Text       string
	RegexpName string
	Regexp     *regexp.Regexp
}

func parse(inputLine string) *line {
	prefix, text := separate(inputLine)

	line := &line{
		Prefix: prefix,
		Text:   text,
	}

	if !line.valid() {
		panic("lines must begin with <, > or #")
	}

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

func (line *line) valid() bool {
	return line.isRequest() || line.isResponse() || line.isComment()
}

func (line *line) isRequest() bool {
	return strings.HasPrefix(string(line.Prefix[0]), ">")
}

func (line *line) isResponse() bool {
	return strings.HasPrefix(string(line.Prefix[0]), "<")
}

func (line *line) isComment() bool {
	return strings.HasPrefix(string(line.Prefix[0]), "#")
}

func (line *line) String() string {
	return line.Text
}
