package main

type specTriplet struct {
	DesiredRequest   *request
	ExpectedResponse *response
	ActualResponse   *response
}

func (specTriplet *specTriplet) isRequestOnly() bool {
	return specTriplet.ExpectedResponse == nil
}
