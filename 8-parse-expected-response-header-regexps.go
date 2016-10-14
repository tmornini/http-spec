package main

import (
	"fmt"
	"strings"
)

func parseExpectedResponseHeaderRegexps(context *context) {
	for _, headerLine := range *context.ExpectedResponse.HeaderLines {
		parts := strings.Split(headerLine.Input, "Δ")

		fmt.Printf("-> %v\n", parts)
	}

	// for bodyLine := range *context.ExpectedResponse.BodyLines {
	// 	parts := strings.Split(headerLine.Input, "Δ")
	//
	// 	fmt.Printf("--> %v", parts)
	// }

	fmt.Println(context.ExpectedResponse)
}
