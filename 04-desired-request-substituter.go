package main

func desiredRequestSubstituter(context *context) {
	context.log("04 desired-request-substituter")

	context.SpecTriplet.DesiredRequest.substitute(context)

	if context.Err != nil {
		return
	}

	if context.SpecTriplet.isRequestOnly() {
		desiredRequestSender(context)

		return
	}

	expectedResponseMatchParser(context)
}
