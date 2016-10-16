package main

import (
	"regexp"
	"strings"
)

func parseExpectedLineMatches(context *context) {
	for _, line := range context.ExpectedResponse.Lines {
		parts := strings.Split(line.Text, "⬚")

		switch len(parts) {
		case 1:
			continue
		case 4:
			line.RegexpName = parts[1]
			line.Regexp = regexp.MustCompile(parts[2])
		default:
			panic("regexp must be in the form ⬚name⬚regexp⬚")
		}
	}

	parseExpectedLineSubstitutions(context)
}
