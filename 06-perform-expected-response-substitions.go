package main

func performExpectedResponseSubstitutions(context *context) error {
	context.log("06 perform-expected-response-substitutions")

	expectedResponse := context.SpecTriplet.ExpectedResponse

	for _, expectedResponseHeaderLine := range expectedResponse.HeaderLines {
		expectedResponseHeaderLine.substitute(context)
	}

	for _, expectedResponseBodyLine := range expectedResponse.BodyLines {
		expectedResponseBodyLine.substitute(context)
	}

	return sendRequestReceiveResponse(context)
}
