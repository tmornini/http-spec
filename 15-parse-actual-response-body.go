package main

import (
	"bufio"
	"fmt"
)

func parseActualBody(context *context) bool {
	context.log("15 parseActualHeaders")

	scanner := bufio.NewScanner(context.HTTPResponse.Body)

	for scanner.Scan() {
		exitWithStatusOneIf(scanner.Err())

		context.ActualResponseLineNumber++

		position :=
			fmt.Sprintf("actual-response:%v", context.ActualResponseLineNumber)

		bodyLine := parse(position, "< "+scanner.Text())

		context.ActualResponse.Lines =
			append(context.ActualResponse.Lines, bodyLine)
	}

	context.HTTPResponse.Body.Close()

	return compareActualToExpected(context)
}
