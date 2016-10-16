package main

import "fmt"

func parseExpectedLineSubstitutions(context *context) {
	for _, line := range context.ExpectedResponse.Lines {
		line.validate()
	}

	fmt.Println(context.Request)
	fmt.Println(context.ExpectedResponse)
}
