package main

import (
	"io"
	"time"
)

func specTripletIterator(context *context) {
	context.enterStage("03 spec-triplet-iterator")

	for {
		context.Err = nil

		desiredRequest, err := requestFromFile(context)

		if errorHandler(context, err) {
			if err != io.EOF {
				context.ResultGathererChannel <- *context
			}

			return
		}

		expectedResponse, err := responseFromFile(context)

		if errorNotEOFHandler(context, err) {
			context.ResultGathererChannel <- *context

			return
		}

		context.SpecTriplet = &specTriplet{
			DesiredRequest:   desiredRequest,
			ExpectedResponse: expectedResponse,
			RequestOnly:      expectedResponse == nil,
		}

		desiredRequestSubstituter(context)

		context.SpecTriplet.Duration = time.Since(context.SpecTriplet.StartedAt)
		context.ResultGathererChannel <- *context
	}
}
