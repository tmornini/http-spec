package main

import (
	"fmt"
	"log"
)

func parseRequestBody(context *context) {
	bodyLines := []string{}

	for context.Scanner.Scan() {
		potentialBodyLine := context.Scanner.Text()

		if potentialBodyLine == "⇓" {
			break
		}

		bodyLines = append(bodyLines, potentialBodyLine)
	}

	context.Request.BodyLines = &bodyLines

	if err := context.Scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(context.Request)

	parseExpectedResponseStatusLine(context)
}
