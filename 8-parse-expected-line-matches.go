package main

import (
	"fmt"
	"strings"
)

func parseExpectedLineMatches(context *context) {
	for _, line := range context.ExpectedResponse.Lines {
		parts := strings.Split(line.Text, "⬚")

		// TODO convert to regexp
		fmt.Println(len(parts))
	}

	parseExpectedLineSubstitutions(context)
}
