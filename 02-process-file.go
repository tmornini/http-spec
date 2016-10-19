package main

import (
	"bufio"
	"os"
)

func processFile(context context) {
	context.log("02 processFile")

	defer context.WaitGroup.Done()

	osFile, err := os.Open(context.Pathname)

	exitWithStatusOneIf(err)

	defer osFile.Close()

	context.Substitutions = map[string]string{}

	context.Scanner = bufio.NewScanner(osFile)

	for {
		eof := parseRequestLine(&context)

		if eof {
			break
		}
	}
}
