package main

import (
	"fmt"
	"regexp"
	"strings"
)

type line struct {
	LineNumber int
	PathName   string
	InputText  string
	IOPrefix   string
	Text       string
	RegexpName string
	Regexp     *regexp.Regexp
}

func lineFromFile(context *context) (*line, error) {
	inputText, err := context.File.readLine()

	if err != nil {
		return nil, err
	}

	line, err :=
		lineFromText(context.File.PathName, context.File.LineNumber, inputText)

	if err != nil {
		return nil, err
	}

	return line, nil
}

func lineFromText(
	pathName string,
	lineNumber int,
	inputText string) (*line, error) {
	ioPrefix, text := split(inputText)

	line := &line{
		PathName:   pathName,
		LineNumber: lineNumber,
		InputText:  inputText,
		IOPrefix:   ioPrefix,
		Text:       text,
	}

	err := line.validate()

	if err != nil {
		return nil, err
	}

	return line, nil

}

func split(inputText string) (string, string) {
	length := len(inputText)

	switch length {
	case 0:
		return "", ""
	case 1:
		return string(inputText[0]), ""
	case 2:
		return string(inputText[0:1]), ""
	default:
		return string(inputText[0:1]), string(inputText[2:length])
	}
}

func (line *line) validate() error {
	if line.isBlank() ||
		line.isEmpty() ||
		line.isRequest() ||
		line.isResponse() ||
		line.isComment() {
		return nil
	}

	return fmt.Errorf("malformed line: %s", line.String())
}

func (line *line) isBlank() bool {
	return line.InputText == ""
}

func (line *line) isEmpty() bool {
	return line.InputText != "" && line.Text == ""
}

func (line *line) isRequest() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), ">")
}

func (line *line) isResponse() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), "<")
}

func (line *line) isComment() bool {
	return line.IOPrefix != "" &&
		strings.HasPrefix(string(line.IOPrefix[0]), "#")
}

func (line *line) substitute(context *context) error {
	parts := strings.Split(line.Text, substitionIdentifier)

	switch len(parts) {
	case 1:
		return nil
	case 3:
		substitution := context.Substitutions[parts[1]]

		if substitution == "" {
			return fmt.Errorf("unknown tag: %v", parts[1])
		}

		line.Text = parts[0] + substitution + parts[2]
	default:
		return fmt.Errorf("malformed substition: %s", line)
	}

	return nil
}

func (line *line) compare(context *context, otherLine *line) error {
	if line.Regexp == nil && line.Text == otherLine.Text {
		return nil
	}

	if line.Regexp == nil && line.Text != otherLine.Text {
		return fmt.Errorf("%v != %v", otherLine, line)
	}

	if line.Regexp != nil {
		matches := line.Regexp.FindStringSubmatch(otherLine.Text)

		if len(matches) == 0 {
			return fmt.Errorf("%v !~ %v", otherLine, line)
		}

		if line.RegexpName != "" {
			context.Substitutions[line.RegexpName] = matches[1]
		}
	}

	return nil
}

func (line *line) String() string {
	return fmt.Sprintf("[%s:%3d] %s %s", line.PathName, line.LineNumber, line.IOPrefix, line.Text)
}
