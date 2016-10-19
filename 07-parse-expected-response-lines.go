package main

func parseExpectedResponseLines(context *context) bool {
	context.log("07 parseExpectedResponseLines")

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
