package main

func parseExpectedResponseLines(context *context) bool {
	context.log("07 parseExpectedResponseLines")

	for {
		responseLine := context.File.readLine()

		if len(responseLine.InputLine) == 0 {
			break
		}

		context.ExpectedResponse.Lines =
			append(context.ExpectedResponse.Lines, responseLine)
	}

	return parseMatches(context)
}
