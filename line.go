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
	ioPrefix, text := split(inputLine)

	line := &line{
		IOPrefix: ioPrefix,
		Text:     text,
	}

	validate(line)

	return line
}

func split(inputLine string) (string, string) {
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

func validate(line *line) {
	if line.isRequest() || line.isResponse() || line.isComment() {
		return
	}

	exitWithStatusOne("lines must begin with <, > or #")
}

func (line *line) substitute(context *context) {
	parts := strings.Split(line.Text, substitionIdentifier)

	switch len(parts) {
	case 1:
		return
	case 3:
		substitution := context.Substitutions[parts[1]]

		if substitution == "" {
			exitWithStatusOne("unknown tag: " + parts[1])
		}

		line.Text = parts[0] + substitution + parts[2]
	default:
		exitWithStatusOne(
			fmt.Sprintf(
				"substition must be formed %scapture-name%s",
				substitionIdentifier,
				substitionIdentifier,
			),
		)
	}
}

func (line *line) compare(context *context, otherLine *line) {
	if line.Regexp == nil {
		if line.Text != otherLine.Text {
			exitWithStatusOne(fmt.Sprintf("%s != %s", line.Text, otherLine.Text))
		}
	} else {
		matches := line.Regexp.FindStringSubmatch(otherLine.Text)

		if len(matches) == 0 {
			exitWithStatusOne(fmt.Sprintf("%s !~ %s", line.Regexp, otherLine))
		} else {
			if line.RegexpName != "" {
				context.Substitutions[line.RegexpName] = matches[1]
			}
		}
	}
}
