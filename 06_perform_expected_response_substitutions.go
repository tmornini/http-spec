package main

func performExpectedResponseSubstitutions(state *state) error {
	state.log("perform-expected-response-substitutions")

	state.SpecTriplet.ExpectedResponse.substitute(state)

	return sendRequest(state)
}
