package main

import "fmt"

func parseActualBlankLine(context *context) bool {
	context.log("14 parseBlankLine")

	context.ActualResponseLineNumber++

	position :=
		fmt.Sprintf("actual-response:%v", context.ActualResponseLineNumber)

	blankLine := parse(position, "< ")

	context.ActualResponse.Lines =
		append(context.ActualResponse.Lines, blankLine)

	return parseActualBody(context)
}
