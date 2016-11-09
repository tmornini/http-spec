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

	startedAt := time.Now()

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
					"%ssuccess %s %s%s\n",
					Green,
					completedContext.SpecTriplet.Duration.String(),
					completedContext.SpecTriplet.String(),
					Reset,
				)
		} else {
			success = false
			failureCount++

			location := "failure "

			if completedContext.File == nil {
				location += "[" + completedContext.Pathname + "]"
			} else {
				if completedContext.SpecTriplet == nil {
					location += completedContext.File.String()
				} else {
					location +=
						completedContext.SpecTriplet.Duration.String() + " " +
							completedContext.SpecTriplet.String()

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
					location,
					completedContext.Err.Error(),
					Reset,
				)
		}
	}

	duration := time.Since(startedAt)

	fmt.Println()

	for _, result := range outputs {
		fmt.Print(result)
	}

	if successCount != 0 {
		fmt.Printf("%ssuccess count: %d%s\n", Green, successCount, Reset)
	}

	if failureCount != 0 {
		fmt.Printf("%sfailure count: %d%s\n", Red, failureCount, Reset)
	}

	if !success {
		fmt.Printf("%sFAILURE: %s%s\n", Red, duration.String(), Reset)
		os.Exit(1)
	}

	fmt.Printf("%sSUCCESS: %s%s\n", Green, duration.String(), Reset)
	os.Exit(0)
}
