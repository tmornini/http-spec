package main

import (
	"fmt"
	"regexp"
	"strings"
)

func expectedResponseMatchParser(context *context) {
	context.log("05 expected-response-match-parser")

	expectedResponse := context.SpecTriplet.ExpectedResponse

	for _, line := range expectedResponse.allHeaderAndBodyLines() {
		parts := strings.Split(line.Text, regexpIdentifier)

		count := len(parts)

		if count == 1 {
			continue
		}

		if count == 0 || count == 2 || count == 3 || (count-4)%3 != 0 {
			errorHandler(
				context,
				fmt.Errorf(
					"regexp must be formed %soptional-capture-name%sregexp%s",
					regexpIdentifier,
					regexpIdentifier,
					regexpIdentifier,
				),
			)

			return
		}

		reString := regexp.QuoteMeta(parts[0])

		for i := 1; i < count-1; i += 3 {
			line.RegexpNames = append(line.RegexpNames, parts[i])

			reString +=
				"(" +
					parts[i+1] +
					")" +
					regexp.QuoteMeta(parts[i+2])
		}

		regexp, err := regexp.Compile(reString)

		if errorHandler(context, err) {
			return
		}

		line.Regexp = regexp
	}

	expectedResponseSubstituter(context)
}
