package main

func parseExpectedResponseHeaders(context *context) bool {
	context.log("06 parseExpectedResponseHeaders")

	for context.Scanner.Scan() {
		exitWithStatusOneIf(context.Scanner.Err())

		potentialInputLine := context.Scanner.Text()

		if len(potentialInputLine) == 0 {
			break
		}

		line := parse(potentialInputLine)

		context.ExpectedResponse.Lines = append(context.ExpectedResponse.Lines, line)
	}

	return parseMatches(context)
}
