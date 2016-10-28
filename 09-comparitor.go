package main

import "fmt"

func responseComparitor(context *context) {
	context.log("09 comparitor")

	if context.SpecTriplet.ExpectedResponse == nil {
		fmt.Println(context.SpecTriplet.ActualResponse)

		errorHandler(
			context,
			fmt.Errorf("no expected response"),
		)

		return
	}

	expectedResponse := context.SpecTriplet.ExpectedResponse
	actualResponse := context.SpecTriplet.ActualResponse

	expectedResponseLines := expectedResponse.allHeaderAndBodyLines()
	actualResponseLines := actualResponse.allHeaderAndBodyLines()

	if len(actualResponseLines) != len(expectedResponseLines) {
		errorHandler(
			context,
			fmt.Errorf(
				"response line count(%d) differs from expected line count (%d)",
				len(actualResponseLines),
				len(expectedResponseLines),
			),
		)

		return
	}

	for i, expectedResponseLine := range expectedResponseLines {
		err := expectedResponseLine.compare(context, actualResponseLines[i])

		if errorHandler(context, err) {
			return
		}
	}
}
