package main

import (
	"io"
	"time"
)

func specTripletIterator(context *context) {
	context.enterStage("03 spec-triplet-iterator")

	for {
		desiredRequest, err := requestFromFile(context)

		if err == io.EOF {
			return
		}

		if err != nil {
			context.Err = err
			context.ResultGathererChannel <- *context

			return
		}

		expectedResponse, err := responseFromFile(context)

		if err != nil && err != io.EOF {
			context.Err = err
			context.ResultGathererChannel <- *context

			return
		}

		tripletContext := *context
		tripletContext.SpecTriplet = &specTriplet{
			DesiredRequest:   desiredRequest,
			ExpectedResponse: expectedResponse,
			RequestOnly:      expectedResponse == nil,
		}

		desiredRequestSubstituter(&tripletContext)

		tripletContext.SpecTriplet.Duration =
			time.Since(tripletContext.SpecTriplet.StartedAt)
		context.ResultGathererChannel <- tripletContext
	}
}
