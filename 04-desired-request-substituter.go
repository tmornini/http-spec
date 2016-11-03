package main

func desiredRequestSubstitor(context *context) {
	context.log("04 desired-request-substituter")

	context.SpecTriplet.DesiredRequest.substitute(context)

	if context.SpecTriplet.ExpectedResponse == nil {
		desiredRequestSender(context)

		return
	}

	expectedResponseMatchParser(context)
}
