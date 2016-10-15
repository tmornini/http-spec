package main

func parseRequestLines(context *context) {
	var lines []*line

	for context.Scanner.Scan() {
		potentialLine := context.Scanner.Text()

		if potentialLine == "" {
			break
		}

		line := parse(potentialLine)

		lines = append(lines, line)
	}

	if err := context.Scanner.Err(); err != nil {
		panic(err)
	}

	context.Request.Lines = lines

	parseExpectedStatusLine(context)
}
