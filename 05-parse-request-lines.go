package main

func parseRequestLines(context *context) bool {
	context.log("05 parseRequestLines")

	for context.Scanner.Scan() {
		exitWithStatusOneIf(context.Scanner.Err())

		potentialLine := context.Scanner.Text()

		if len(potentialLine) == 0 {
			break
		}

		line := parse(potentialLine)

		context.Request.Lines = append(context.Request.Lines, line)
	}

	return parseExpectedStatusLine(context)
}
