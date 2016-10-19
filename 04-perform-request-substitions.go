package main

func performRequestSubstitutions(context *context) error {
	context.log("04 perform-request-substitutions")

	context.SpecTriplet.DesiredRequest.substitute(context)

	if context.SpecTriplet.ExpectedResponse == nil {
		return sendRequestReceiveResponse(context)
	}

	return parseExpectedResponseRegexps(context)
}
