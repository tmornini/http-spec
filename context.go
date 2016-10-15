package main

import "bufio"

type context struct {
	Prefix           string
	Pathname         string
	Scanner          *bufio.Scanner
	Request          *request
	ExpectedResponse *response
	ActualResponse   *response
}
