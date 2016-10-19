package main

func compareActualToExpected(context *context) bool {
	context.log("11 compareActualToExpected")

	context.ActualResponse.compare(context, context.ExpectedResponse)

	return false
}
