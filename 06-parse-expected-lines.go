package main

func parseExpectedResponseHeaders(context *context) bool {
	context.log("06 parseExpectedResponseHeaders")

	lines := []*line{}

	for context.Scanner.Scan() {
		exitWithStatusOneIf(context.Scanner.Err())

		potentialInputLine := context.Scanner.Text()

		if len(potentialInputLine) == 0 {
			break
		}

		line := parse(potentialInputLine)

		lines = append(lines, line)
	}

	context.ExpectedResponse.Lines = lines

	return parseMatches(context)
}
