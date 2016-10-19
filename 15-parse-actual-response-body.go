package main

import "bufio"

func parseActualBody(context *context) bool {
	context.log("14 parseActualHeaders")

	scanner := bufio.NewScanner(context.HTTPResponse.Body)

	for scanner.Scan() {
		exitWithStatusOneIf(scanner.Err())

		line := parse("< " + scanner.Text())

		context.ActualResponse.Lines = append(context.ActualResponse.Lines, line)
	}

	context.HTTPResponse.Body.Close()

	return compareActualToExpected(context)
}
