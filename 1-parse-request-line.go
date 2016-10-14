package main

import (
	"fmt"
	"strings"
)

func parseRequestLine(context *context) {
	context.Scanner.Scan()

	requestLine := context.Scanner.Text()

	words := strings.Split(requestLine, " ")

	context.Verb = words[0]
	context.Path = words[1]
	context.Version = words[2]

	fmt.Printf("%#v\n", context)
}
