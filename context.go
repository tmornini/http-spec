package main

import "bufio"

type context struct {
	UriScheme        string
	Pathname         string
	Scanner          *bufio.Scanner
	Request          *request
	ExpectedResponse *response
	ActualResponse   *response
	Captures         *map[string]string
}
