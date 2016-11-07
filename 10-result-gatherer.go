package main

import (
	"fmt"
	"os"
	"time"
)

func resultGatherer(context context) {
	context.log("10 result-gatherer")

	success := true

	successCount := 0
	failureCount := 0

	startedAt := time.Now()

	for completedContext := range context.ResultGathererChannel {
		if completedContext.Err == nil &&
			completedContext.SpecTriplet != nil &&
			completedContext.SpecTriplet.isRequestOnly() {
			fmt.Println(completedContext.SpecTriplet.ActualResponse)

			completedContext.Err = fmt.Errorf("no expected response")
		}

		if completedContext.Err == nil {
			successCount++
			fmt.Printf(
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
				}
			}

			fmt.Fprintf(
				os.Stderr,
				"%s%s %s%s\n",
				Red,
				location,
				completedContext.Err.Error(),
				Reset,
			)
		}
	}

	duration := time.Since(startedAt)

	fmt.Printf("%sTotal successes: %d%s\n", Green, successCount, Reset)
	fmt.Printf("%sTotal failures: %d%s\n", Red, failureCount, Reset)
	fmt.Printf("%sTotal run time: %s%s\n", Reset, duration.String(), Reset)

	if success {
		os.Exit(0)
	}

	os.Exit(1)
}
