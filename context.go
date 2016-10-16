package main

import (
	"bufio"
	"fmt"
)

type context struct {
	URIScheme        string
	Pathname         string
	Scanner          *bufio.Scanner
	Request          *request
	ExpectedResponse *response
	ActualResponse   *response
	Substitutions    map[string]string
}

func (context *context) log(tag string) {
	if logFunctions {
		fmt.Println(tag)
	}

	if logContext {
		fmt.Printf("%#v\n", context)
	}
}
