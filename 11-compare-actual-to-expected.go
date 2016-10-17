package main

func compareActualToExpected(context *context) bool {
	context.log("11 compareActualToExpected")

	for i, line := range context.ExpectedResponse.Lines {
		context.ActualLine = context.ActualResponse.Lines[i].Text

		line.compare(context)
	}

	return false
}
