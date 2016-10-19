package main

func compareActualToExpected(context *context) bool {
	context.log("12 compareActualToExpected")

	context.ExpectedResponse.compare(context, context.ActualResponse)

	return false
}
