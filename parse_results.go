package main

import (
	"fmt"
)

func parseResults(in <-chan state) <-chan state {
	out := make(chan state)

	go func() {
		// successCount := 0
		// failureCount := 0
		// outputs := make(map[string]string)
		// startedAt := time.Now()

		for s := range in {
			s.log("parse-results")

			if s.Error != nil {
				// Failure
				// failureCount++
				s.Result.Failure = true

				location := ""
				response := ""

				if s.File == nil {
					// file open failure
					location += "[" + s.HTSFFilePath + "]"
				} else {
					if s.SpecTriplet == nil {
						// request/response parsing failure
						location += s.File.String()
					} else {
						// request/response matching failure
						location +=
							s.SpecTriplet.String() +
								" " +
								s.SpecTriplet.Duration.String()

						if s.SpecTriplet.ActualResponse != nil {
							response = s.SpecTriplet.ActualResponse.String() + "\n"
						}
					}
				}

				// outputs[s.ID] +=
				output :=
					fmt.Sprintf(
						"%sSubstitutions: %#v%s\n",
						Yellow,
						s.Substitutions,
						Reset,
					)

				// outputs[s.ID] +=
				output +=
					fmt.Sprintf(
						"%s%s%s %s\n%s\n",
						Red,
						location,
						Reset,
						s.Error.Error(),
						response,
					)

				s.Result.Message = output

				out <- s

			} else {
				// Success
				// successCount++
				s.Result.Failure = false

				// outputs[s.ID] +=
				s.Result.Message =
					fmt.Sprintf(
						"%s%s%s %s\n",
						Green,
						s.SpecTriplet.String(),
						Reset,
						s.SpecTriplet.Duration.String(),
					)

				out <- s
			}
		}

		// duration := time.Since(startedAt)

		// fmt.Println()

		// for _, result := range outputs {
		// 	fmt.Print(result)
		// }

		// if failureCount > 0 {
		// 	fmt.Printf(
		// 		"%sFAILURE:%s %s+%d%s %s-%d%s %s\n",
		// 		Red,
		// 		Reset,
		// 		Green,
		// 		successCount,
		// 		Reset,
		// 		Red,
		// 		failureCount,
		// 		Reset,
		// 		duration.String(),
		// 	)
		//
		// 	// os.Exit(1)
		// 	out <- 1
		//
		// } else {
		// 	fmt.Printf(
		// 		"\n%sSUCCESS:%s %s+%d%s %s\n",
		// 		Green,
		// 		Reset,
		// 		Green,
		// 		successCount,
		// 		Reset,
		// 		duration.String(),
		// 	)
		// }
		//
		// // os.Exit(0)
		// out <- 0

		close(out)
	}()

	return out
}
