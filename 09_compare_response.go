package main

import (
	"fmt"
	"time"
)

func compareResponse(state *state) error {
	state.log("compare-response")

	expectedResponseLines := state.SpecTriplet.ExpectedResponse.allLines()
	actualResponseLines := state.SpecTriplet.ActualResponse.allLines()

	if len(expectedResponseLines) != len(actualResponseLines) {
		return fmt.Errorf(
			"expected line count (%d) differs from actual line count(%d)",
			len(expectedResponseLines),
			len(actualResponseLines),
		)
	}

	for i, expectedResponseLine := range expectedResponseLines {
		err := expectedResponseLine.compare(state, actualResponseLines[i])

		if err != nil {
			return err
		}
	}

	time.Sleep(state.SpecTriplet.ExpectedResponse.Duration)

	return nil
}
