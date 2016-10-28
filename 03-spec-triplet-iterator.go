package main

func specTripletIterator(context *context) {
	context.log("03 spec-triplet-iterator")

	for {
		context.Err = nil

		desiredRequest, err := requestFromFile(context)

		if errorHandler(context, err) {
			return
		}

		expectedResponse, err := responseFromFile(context)

		if errorNotEOFHandler(context, err) {
			return
		}

		context.SpecTriplet = &specTriplet{
			DesiredRequest:   desiredRequest,
			ExpectedResponse: expectedResponse,
		}

		desiredRequestSubstitor(context)

		context.ResultGathererChannel <- *context
	}
}
