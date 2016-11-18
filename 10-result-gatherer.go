package main

import (
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"
)

func resultGatherer(context context) {
	context.log("10 result-gatherer")

	success := true

	successCount := 0
	failureCount := 0

	outputs := map[*big.Int]string{}

	for completedContext := range context.ResultGathererChannel {
		// if completedContext.Err == nil &&
		// 	completedContext.SpecTriplet != nil &&
		// 	completedContext.SpecTriplet.isRequestOnly() {
		// 	outputs[completedContext.ID] +=
		// 		completedContext.SpecTriplet.ActualResponse.String() + "\n"
		// 	completedContext.Err = fmt.Errorf("no expected response")
		// }

		if completedContext.Err == nil {
			successCount++
			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s %s%s\n",
					Green,
					completedContext.SpecTriplet.String(),
					completedContext.SpecTriplet.Duration.String(),
					Reset,
				)
		} else {
			success = false
			failureCount++

			output := ""

			if completedContext.File == nil {
				output += "[" + completedContext.Pathname + "]"
			} else {
				if completedContext.SpecTriplet == nil {
					output += completedContext.File.String()
				} else {
					output +=
						completedContext.SpecTriplet.String() + " " +
							completedContext.SpecTriplet.Duration.String()

					if completedContext.SpecTriplet.isRequestOnly() ||
						strings.HasPrefix(
							completedContext.Err.Error(), "response line count") {
						outputs[completedContext.ID] +=
							completedContext.SpecTriplet.ActualResponse.String() + "\n"
					}
				}
			}

			outputs[completedContext.ID] +=
				fmt.Sprintf(
					"%s%s %s%s\n",
					Red,
					output,
					completedContext.Err.Error(),
					Reset,
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
			"%sFAILURE:%s %s+%d%s %s-%d%s in %s\n",
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
		"%sSUCCESS:%s %s+%d%s in %s\n",
		Green,
		Reset,
		Green,
		successCount,
		Reset,
		duration.String(),
	)

	os.Exit(0)
}
