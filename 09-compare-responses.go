package main

import "fmt"

func compareResponses(context *context) error {
	context.log("09 compare-responses")

	expectedResponse := context.SpecTriplet.ExpectedResponse
	actualResponse := context.SpecTriplet.ActualResponse

	expectedResponseLines := expectedResponse.allHeaderAndBodyLines()
	actualResponseLines := actualResponse.allHeaderAndBodyLines()

	if len(actualResponseLines) != len(expectedResponseLines) {
		return fmt.Errorf(
			"response line count(%d) differs from expected line count (%d)",
			len(actualResponseLines),
			len(expectedResponseLines),
		)
	}

	for i, expectedResponseLine := range expectedResponseLines {
		err := expectedResponseLine.compare(context, actualResponseLines[i])

		if err != nil {
			return err
		}
	}

	// report per file status

	return nil
}
