package main

import "io"

func parseSpecTriplets(context *context) error {
	context.log("03 parse-spec-triplets")

	for {
		desiredRequest, err := requestFromFile(context)

		if err != nil {
			return err
		}

		expectedResponse, err := responseFromFile(context)

		if err != nil && err != io.EOF {
			return err
		}

		context.SpecTriplet = &specTriplet{
			DesiredRequest:   desiredRequest,
			ExpectedResponse: expectedResponse,
		}

		return performRequestSubstitutions(context)
	}
}
