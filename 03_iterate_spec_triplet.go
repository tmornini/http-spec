package main

import (
	"io"
	"time"
)

func iterateSpecTriplet(in <-chan state) <-chan state {
	out := make(chan state)

	go func() {
		for s := range in {
			s.log("iterate-spec-triplet")

			// Handle any errors
			if s.Error != nil {
				out <- s
				continue
			}

			// Loop to handle multiple tests per file
		ForLoop:
			for {
				s.log("iterate-spec-triplet-for-loop")

				desiredRequest, err := getRequestFromFile(&s)

				if err != nil {
					if err == io.EOF {
						// Finished with the file
						break ForLoop
					}

					s.Error = err
					out <- s

					break ForLoop
				}

				expectedResponse, err := getResponseFromFile(&s)

				if err != nil && err != io.EOF {
					s.Error = err
					out <- s

					break ForLoop
				}

				s.SpecTriplet = &specTriplet{
					DesiredRequest:   desiredRequest,
					ExpectedResponse: expectedResponse,
				}

				err = performRequestSubstitutions(&s)

				if err != nil {
					s.Error = err
					out <- s

					break ForLoop
				}

				s.SpecTriplet.Duration = time.Since(s.SpecTriplet.StartedAt)

				out <- s
			}

			// Close the file after we have read all the lines
			s.File.OSFile.Close()
		}

		close(out)
	}()

	return out
}
