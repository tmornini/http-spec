package main

import (
	"fmt"
	"os"
)

const regexpIdentifier = "⧆"
const substitionIdentifier = "⧈"

const logFunctions = false
const logContext = false

func main() {
	if logFunctions {
		fmt.Println("00 main")
	}

	processCLIArguments()
}

func exitWithStatusOne(message string) {
	os.Stderr.WriteString(message + "\n")
	os.Exit(1)
}

func exitWithStatusOneIf(err error) {
	if err != nil {
		exitWithStatusOne(err.Error())
	}
}
