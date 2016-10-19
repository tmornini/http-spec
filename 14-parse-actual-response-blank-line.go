package main

func parseActualBlankLine(context *context) bool {
	context.log("13 parseBlankLine")

	blankLine := parse("< ")

	context.ActualResponse.Lines =
		append(context.ActualResponse.Lines, blankLine)

	return parseActualBody(context)
}
