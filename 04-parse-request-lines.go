package main

func parseRequestLines(context *context) bool {
	context.log("4 parseRequestLines")

	var lines []*line

	for context.Scanner.Scan() {
		panicOn(context.Scanner.Err())

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
