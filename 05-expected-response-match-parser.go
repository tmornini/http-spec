package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func expectedResponseMatchParser(context *context) {
	context.log("05 expected-response-match-parser")

	expectedResponse := context.SpecTriplet.ExpectedResponse

	for _, line := range expectedResponse.allLines() {
		parts := strings.Split(line.Text, regexpIdentifier)

		count := len(parts)

		if count == 1 {
			continue
		}

		if (count-4)%3 != 0 {
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

		var re *regexp.Regexp
		var err error

		if parts[0] != "" {
			re, err = regexp.Compile(regexp.QuoteMeta(parts[0]))

			if errorHandler(context, err) {
				return
			}

			line.RegexpNames = append(line.RegexpNames, ":prefix")
			line.Regexps = append(line.Regexps, re)
		}

		for i := 1; i < count-1; i += 3 {
			captureName := parts[i]

			if captureName == ":prefix" || captureName == ":postfix" {
				errorHandler(
					context,
					errors.New("capture names cannot be :prefix or :postfix"),
				)

				return
			}

			reString := "("

			_, ok := context.Matchers[parts[i+1]]

			if ok {
				reString += context.Matchers[parts[i+1]]
			} else {
				reString += parts[i+1]
			}

			reString += ")"

			re, err = regexp.Compile(reString)

			if errorHandler(context, err) {
				return
			}

			line.RegexpNames = append(line.RegexpNames, parts[i])
			line.Regexps = append(line.Regexps, re)

			postfix := parts[i+2]

			if postfix != "" {
				re, err = regexp.Compile(regexp.QuoteMeta(postfix))

				if errorHandler(context, err) {
					return
				}

				line.RegexpNames = append(line.RegexpNames, ":postfix")
				line.Regexps = append(line.Regexps, re)
			}
		}
	}

	expectedResponseSubstituter(context)
}
