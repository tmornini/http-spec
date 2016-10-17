package main

import (
	"bufio"
	"os"
)

func processFile(context *context) {
	context.log("02 processFile")

	osFile, err := os.Open(context.Pathname)

	panicOn(err)

	defer osFile.Close()

	context.Substitutions = map[string]string{}

	context.Scanner = bufio.NewScanner(osFile)

	for {
		eof := parseRequestLine(context)

		if eof {
			break
		}
	}
}
