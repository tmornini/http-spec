package main

import (
	"fmt"
	"regexp"
	"strings"
)

type line struct {
	LineNumber  int
	PathName    string
	InputText   string
	IOPrefix    string
	Text        string
	RegexpNames []string
	Regexps     []*regexp.Regexp
}

func getNextLineFromFile(state *state) (*line, error) {
	var inputText string
	var err error
	var line *line

	for {
		inputText, err = state.File.readLine()

		if err != nil {
			return nil, err
		}

		line, err = newLineFromText(
			state.File.PathName,
			state.File.LineNumber,
			inputText,
		)

		if err != nil {
			return nil, err
		}

		if !line.isComment() {
			break
		}
	}

	return line, nil
}

func newLineFromText(pathName string, lineNumber int, inputText string) (*line, error) {
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
		line.isComment() ||
		line.isEmpty() ||
		line.isRequest() ||
		line.isResponse() ||
		line.isSleep() {
		return nil
	}

	return fmt.Errorf("malformed line: %s", line.String())
}

func (line *line) isBlank() bool {
	return line.InputText == ""
}

func (line *line) isComment() bool {
	return line.IOPrefix != "" && string(line.IOPrefix[0]) == "#"
}

func (line *line) isEmpty() bool {
	return line.InputText != "" && line.Text == ""
}

func (line *line) isSleep() bool {
	return line.IOPrefix != "" && string(line.IOPrefix[0]) == "+"
}

func (line *line) isRequest() bool {
	return line.IOPrefix != "" && string(line.IOPrefix[0]) == ">"
}

func (line *line) isResponse() bool {
	return line.IOPrefix != "" && string(line.IOPrefix[0]) == "<"
}

func (line *line) substitute(state *state) error {
	parts := strings.Split(line.Text, substitutionIdentifier)

	count := len(parts)

	if count == 1 {
		return nil
	}

	if (count-3)%2 != 0 {
		return fmt.Errorf("malformed substitution: %s", line)
	}

	substitutedText := parts[0]

	for i := 1; i < count-1; i += 2 {
		substitution, known := state.Substitutions[parts[i]]

		if !known {
			return fmt.Errorf("unknown tag: %v", parts[i])
		}

		substitutedText += substitution + parts[i+1]
	}

	line.Text = substitutedText

	return nil
}

func comparisonError(line, otherLine *line) error {
	return fmt.Errorf(
		"line: %d: |%v| != |%v|",
		line.LineNumber,
		line.Content(),
		otherLine.Content(),
	)
}

func (line *line) compare(state *state, otherLine *line) error {
	if line.RegexpNames == nil && line.Text == otherLine.Text {
		return nil
	}

	if line.RegexpNames == nil && line.Text != otherLine.Text {
		return comparisonError(line, otherLine)
	}

	for i, regexpName := range line.RegexpNames {
		match := line.Regexps[i].FindString(otherLine.Text)

		if match == "" {
			return comparisonError(line, otherLine)
		}

		if regexpName != "" && regexpName != ":prefix" && regexpName != ":postfix" {
			state.Substitutions[regexpName] = match
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
