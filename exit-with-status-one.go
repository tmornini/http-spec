package main

import "os"

func exitWithStatusOne(message string) {
	os.Stderr.WriteString(message + "\n")
	os.Exit(1)
}

func exitWithStatusOneIf(err error) {
	if err != nil {
		exitWithStatusOne(err.Error())
	}
}
