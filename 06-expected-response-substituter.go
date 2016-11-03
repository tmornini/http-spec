package main

func expectedResponseSubstituter(context *context) {
	context.log("06 expected-response-substituter")

	context.SpecTriplet.ExpectedResponse.substitute(context)

	desiredRequestSender(context)
}
