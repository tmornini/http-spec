package main

import (
	"bufio"
	"os"
)

func parseFile(context *context) {
	osFile, err := os.Open(context.Pathname)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(osFile)

	context.Scanner = scanner

	parseRequestLine(context)

	osFile.Close()
}
