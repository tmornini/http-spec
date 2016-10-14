package main

type headerLine struct {
	Input string
}

func (headerLine *headerLine) String() string {
	return headerLine.Input
}
