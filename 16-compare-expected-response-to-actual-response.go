package main

func compareActualToExpected(context *context) bool {
	context.log("16 compareActualToExpected")

	context.ExpectedResponse.compare(context, context.ActualResponse)

	return false
}
