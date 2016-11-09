package main

import (
	"fmt"
	"regexp"
	"strings"
)

func newLineFromFile(context *context) (*line, error) {
	var inputText string
	var err error
	var line *line

	for {
		inputText, err = context.File.readLine()

		if err != nil {
			return nil, err
		}

		line, err =
			newLineFromText(context.File.PathName, context.File.LineNumber, inputText)

		if err != nil {
			return nil, err
		}

		if !line.isComment() {
			break
		}
	}

	return line, nil
}

func newLineFromText(
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

type line struct {
	LineNumber  int
	PathName    string
	InputText   string
	IOPrefix    string
	Text        string
	RegexpNames []string
	Regexps     []*regexp.Regexp
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
	parts := strings.Split(line.Text, substitutionIdentifier)

	count := len(parts)

	if count == 1 {
		return nil
	}

	if count == 0 || count == 2 || (count-3)%2 != 0 {
		return fmt.Errorf("malformed substitution: %s", line)
	}

	substitutedLine := parts[0]

	for i := 1; i < count-1; i += 2 {
		substitution, ok := context.Substitutions[parts[i]]

		if !ok {
			return fmt.Errorf("unknown tag: %v", parts[i])
		}

		substitutedLine += substitution + parts[i+1]
	}

	line.Text = substitutedLine

	return nil
}

func (line *line) compare(context *context, otherLine *line) error {
	if line.RegexpNames == nil && line.Text == otherLine.Text {
		return nil
	}

	if line.RegexpNames == nil && line.Text != otherLine.Text {
		return fmt.Errorf("%v != %v", line.Content(), otherLine.Content())
	}

	for i, regexpName := range line.RegexpNames {
		match := line.Regexps[i].FindString(otherLine.Text)

		if match == "" {
			return fmt.Errorf("%v !~ %v", line.Content(), otherLine.Content())
		}

		if regexpName != "" {
			context.Substitutions[regexpName] = match
		}
	}

	return nil
}

func (line *line) Location() string {
	return fmt.Sprintf("[%s:%3d]", line.PathName, line.LineNumber)
}

func (line *line) Content() string {
	return fmt.Sprintf("%s %s", line.IOPrefix, line.Text)
}

func (line *line) String() string {
	return fmt.Sprintf("%s %s", line.Location(), line.Content())
}
