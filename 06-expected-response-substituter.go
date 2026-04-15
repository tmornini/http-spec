package main

func expectedResponseSubstituter(context *context) {
	context.enterStage("06 expected-response-substituter")

	context.SpecTriplet.ExpectedResponse.substitute(context)

	if context.Err != nil {
		return
	}

	desiredRequestSender(context)
}
