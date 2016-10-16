package main

import (
	"bufio"
	"os"
)

func processFile(context *context) {
	context.log("2 processFile")

	osFile, err := os.Open(context.Pathname)

	if err != nil {
		panic(err)
	}

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
