package main

func parseRequestLines(context *context) bool {
	context.log("05 parseRequestLines")

	for {
		anotherRequestLine := context.File.readLine()

		if len(anotherRequestLine.InputLine) == 0 {
			break
		}

		context.Request.Lines = append(context.Request.Lines, anotherRequestLine)
	}

	return parseExpectedStatusLine(context)
}
