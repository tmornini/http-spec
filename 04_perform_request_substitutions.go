package main

func performRequestSubstitutions(state *state) error {
	state.log("perform-request-substitutions")

	state.SpecTriplet.DesiredRequest.substitute(state)

	if state.SpecTriplet.isRequestOnly() {
		return sendRequest(state)
	}

	return parseExpectedResponse(state)
}
