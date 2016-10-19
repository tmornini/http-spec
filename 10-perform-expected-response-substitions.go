package main

func performExpectedResponseSubstitutions(context *context) bool {
	context.log("10 performExpectedResponseSubstitutions")

	for _, line := range context.ExpectedResponse.Lines {
		line.substitute(context)
	}

	return makeActualRequest(context)
}
