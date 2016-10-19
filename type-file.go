package main

import (
	"bufio"
	"fmt"
	"os"
)

type file struct {
	Pathname   string
	OSFile     *os.File
	Scanner    *bufio.Scanner
	LineNumber int
}

func newFile(pathname string) *file {
	osFile, err := os.Open(pathname)

	exitWithStatusOneIf(err)

	return &file{
		Pathname: pathname,
		OSFile:   osFile,
		Scanner:  bufio.NewScanner(osFile),
	}
}

func (file *file) readLine() *line {
	file.Scanner.Scan()

	exitWithStatusOneIf(file.Scanner.Err())

	file.LineNumber++

	inputText := file.Scanner.Text()

	return parse(file.String(), inputText)
}

func (file *file) close() {
	file.OSFile.Close()
}

func (file *file) String() string {
	return fmt.Sprintf("%s:%v", file.Pathname, file.LineNumber)
}
