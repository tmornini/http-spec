package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type file struct {
	*bufio.Reader
	PathName   string
	OSFile     *os.File
	LineNumber int
}

func (file *file) readLine() (string, error) {
	inputText, err := file.ReadString(byte('\n'))

	if err != nil {
		return "", err
	}

	file.LineNumber++

	inputText = strings.TrimSuffix(inputText, "\n")
	inputText = strings.TrimSuffix(inputText, "\r")

	return inputText, nil
}

func (file *file) String() string {
	return fmt.Sprintf("[%s:%3d]", file.PathName, file.LineNumber)
}
