package main

import "bufio"

type context struct {
	Pathname string
	Scanner  *bufio.Scanner
	Request  *request
	Response *response
}
