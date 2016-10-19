package main

func parseRequestLines(context *context) bool {
	context.log("04 parseRequestLines")

	var lines []*line

	for context.Scanner.Scan() {
		exitWithStatusOneIf(context.Scanner.Err())

		potentialLine := context.Scanner.Text()

		if len(potentialLine) == 0 {
			break
		}

		line := parse(potentialLine)

		lines = append(lines, line)
	}

	context.Request.Lines = lines

	return parseExpectedStatusLine(context)
}
