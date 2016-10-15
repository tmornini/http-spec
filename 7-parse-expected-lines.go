package main

func parseExpectedResponseHeaders(context *context) {
	lines := []*line{}

	for context.Scanner.Scan() {
		potentialInputLine := context.Scanner.Text()

		if potentialInputLine == "" {
			break
		}

		line := parse(potentialInputLine)

		lines = append(lines, line)
	}

	err := context.Scanner.Err()

	if err != nil {
		panic(err)
	}

	context.ExpectedResponse.Lines = lines

	parseExpectedLineMatches(context)
}
