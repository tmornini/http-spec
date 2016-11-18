package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

func resultGatherer(context context) {
	context.log("10 result-gatherer")

	success := true

	successCount := 0
	failureCount := 0

	outputs := map[*big.Int]string{}

	for completedContext := range context.ResultGathererChannel {
		if completedContext.Err == nil {
			// SUCCESS
			successCount++
			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n",
					Green,
					completedContext.SpecTriplet.String(),
					Reset,
					completedContext.SpecTriplet.Duration.String(),
				)
		} else {
			// FAILURE
			success = false
			failureCount++

			location := ""
			response := ""

			if completedContext.File == nil {
				// file open failure
				location += "[" + completedContext.Pathname + "]"
			} else {
				if completedContext.SpecTriplet == nil {
					// request/response parsing failure
					location += completedContext.File.String()
				} else {
					// request/response matching failure
					location +=
						completedContext.SpecTriplet.String() + " " +
							completedContext.SpecTriplet.Duration.String()

					response = completedContext.SpecTriplet.ActualResponse.String() + "\n"

				}
			}

			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s%s %s\n%s\n",
					Red,
					location,
					Reset,
					completedContext.Err.Error(),
					response,
				)
		}
	}

	duration := time.Since(context.StartedAt)

	fmt.Println()

	for _, result := range outputs {
		fmt.Print(result)
	}

	if !success {
		fmt.Printf(
			"%sFAILURE:%s %s+%d%s %s-%d%s %s\n",
			Red,
			Reset,
			Green,
			successCount,
			Reset,
			Red,
			failureCount,
			Reset,
			duration.String(),
		)

		os.Exit(1)
	}

	fmt.Printf(
		"\n%sSUCCESS:%s %s+%d%s %s\n",
		Green,
		Reset,
		Green,
		successCount,
		Reset,
		duration.String(),
	)

	os.Exit(0)
}
