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

func (f *file) readLine() (string, error) {
	inputText, err := f.ReadString(byte('\n'))

	if err != nil {
		return "", err
	}

	f.LineNumber++

	inputText = strings.TrimSuffix(inputText, "\n")
	inputText = strings.TrimSuffix(inputText, "\r")

	return inputText, nil
}

func (f *file) String() string {
	return fmt.Sprintf("[%s:%3d]", f.PathName, f.LineNumber)
}
