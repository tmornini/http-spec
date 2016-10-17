package main

func parseExpectedResponseHeaders(context *context) bool {
	context.log("6 parseExpectedResponseHeaders")

	lines := []*line{}

	for context.Scanner.Scan() {
		if err := context.Scanner.Err(); err != nil {
			panic(err)
		}

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
