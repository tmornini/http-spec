package main

import (
	"fmt"
	"regexp"
	"strings"
)

func parseMatches(context *context) bool {
	context.log("07 parseMatches")

	for _, line := range context.ExpectedResponse.Lines {
		parts := strings.Split(line.Text, regexpIdentifier)

		switch len(parts) {
		case 1:
			continue
		case 4:
			line.RegexpName = parts[1]

			quotedPrefix := regexp.QuoteMeta(parts[0])
			quotedSuffix := regexp.QuoteMeta(parts[3])

			line.Regexp = regexp.MustCompile(quotedPrefix + parts[2] + quotedSuffix)
		default:
			panic(
				fmt.Sprintf(
					"regexp must be formed %soptional-capture-name%sregexp%s",
					regexpIdentifier,
					regexpIdentifier,
					regexpIdentifier,
				),
			)
		}
	}

	return performSubstitutions(context)
}
