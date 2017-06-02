package main

import (
	"fmt"
	"time"
)

func responseComparitor(context *context) {
	context.log("09 comparitor")

	expectedResponse := context.SpecTriplet.ExpectedResponse
	actualResponse := context.SpecTriplet.ActualResponse

	expectedResponseLines := expectedResponse.allLines()
	actualResponseLines := actualResponse.allLines()

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

	time.Sleep(expectedResponse.Duration)
}
