package main

import "fmt"

func parseExpectedLineSubstitutions(context *context) {
	fmt.Println(context.Request)
	fmt.Println(context.ExpectedResponse)
}
