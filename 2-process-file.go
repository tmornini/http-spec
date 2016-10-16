package main

import (
	"bufio"
	"os"
)

func processFile(context *context) {
	osFile, err := os.Open(context.Pathname)

	if err != nil {
		panic(err)
	}

	context.Captures = &map[string]string{}

	context.Scanner = bufio.NewScanner(osFile)

	parseRequestLine(context)

	osFile.Close()
}
