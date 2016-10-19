package main

import (
	"fmt"
	"regexp"
	"strings"
)

func parseExpectedResponseRegexps(context *context) error {
	context.log("05 parse-expected-response-regexps")

	expectedResponse := context.SpecTriplet.ExpectedResponse

	for _, line := range expectedResponse.allHeaderAndBodyLines() {
		parts := strings.Split(line.Text, regexpIdentifier)

		switch len(parts) {
		case 1:
			continue
		case 4:
			line.RegexpName = parts[1]

			quotedPrefix := regexp.QuoteMeta(parts[0])
			quotedSuffix := regexp.QuoteMeta(parts[3])

			reString := quotedPrefix + "(" + parts[2] + ")" + quotedSuffix

			regexp, err := regexp.Compile(reString)

			if err != nil {
				return err
			}

			line.Regexp = regexp
		default:
			return fmt.Errorf(
				"regexp must be formed %soptional-capture-name%sregexp%s",
				regexpIdentifier,
				regexpIdentifier,
				regexpIdentifier,
			)
		}
	}

	return performExpectedResponseSubstitutions(context)
}
