package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

func parseExpectedResponse(state *state) error {
	state.log("parse-expected-response")

	for _, line := range state.SpecTriplet.ExpectedResponse.allLines() {
		parts := strings.Split(line.Text, regexpIdentifier)

		count := len(parts)

		// line does not contain any substitions
		if count == 1 {
			continue
		}

		// invalid substition format
		if (count-4)%3 != 0 {
			return fmt.Errorf(
				"regexp must be formed %soptional-capture-name%sregexp%s",
				regexpIdentifier,
				regexpIdentifier,
				regexpIdentifier,
			)
		}

		var re *regexp.Regexp
		var err error

		// TODO: when would this ever be empty at this point?
		// This is the beginning of the line and will be tagged :prefix
		if parts[0] != "" {
			re, err = regexp.Compile(regexp.QuoteMeta(parts[0]))

			if err != nil {
				return err
			}

			line.RegexpNames = append(line.RegexpNames, ":prefix")
			line.Regexps = append(line.Regexps, re)
		}

		for i := 1; i < count-1; i += 3 {
			key := parts[i]
			value := parts[i+1]

			// :prefix and :postfix are used as tags for beginning of line and end of line
			if key == ":prefix" || key == ":postfix" {
				return errors.New("capture names cannot be :prefix or :postfix")
			}

			reString := "("

			// Check for any builtin matchers or use what is passed in as default
			switch value {
			case ":date":
				reString +=
					"(Mon|Tue|Wed|Thu|Fri|Sat|Sun), " +
						"(0\\d|1\\d|2\\d|3[01]) " +
						"(Jan|Feb|Mar|Apr|May|Jun|Jul|Aug|Sep|Oct|Nov|Dec) " +
						"2\\d{3} " +
						"(0\\d|1\\d|2[0-3]):" +
						"(0\\d|1\\d|2\\d|3\\d|4\\d|5\\d):" +
						"(0\\d|1\\d|2\\d|3\\d|4\\d|5\\d) " +
						"(A|M|N|Y|Z|UT|GMT|[A-Z]{3}|[+-](0\\d|1[012]))"
			case ":b62:22":
				reString += "[0-9A-Za-z]{22}"
			case ":iso8601:µs:z":
				reString += "\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[.]\\d{6}Z"
			case ":uuid":
				reString +=
					"[[:xdigit:]]{8}-" +
						"[[:xdigit:]]{4}-" +
						"[[:xdigit:]]{4}-" +
						"[[:xdigit:]]{4}-" +
						"[[:xdigit:]]{12}"
			default:
				reString += value
			}

			reString += ")"

			re, err = regexp.Compile(reString)

			if err != nil {
				return err
			}

			line.RegexpNames = append(line.RegexpNames, key)
			line.Regexps = append(line.Regexps, re)

			postfix := parts[i+2]

			if postfix != "" {
				re, err = regexp.Compile(regexp.QuoteMeta(postfix))

				if err != nil {
					return err
				}

				line.RegexpNames = append(line.RegexpNames, ":postfix")
				line.Regexps = append(line.Regexps, re)
			}
		}
	}

	return performExpectedResponseSubstitutions(state)
}
