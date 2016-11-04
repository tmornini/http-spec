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
				"success %s %s\n",
				completedContext.SpecTriplet.Duration.String(),
				completedContext.SpecTriplet.String())
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

			fmt.Println(location, completedContext.Err.Error())
		}
	}

	duration := time.Since(startedAt)

	fmt.Println("Total successes:", successCount)
	fmt.Println("Total failures:", failureCount)
	fmt.Println("Total run time:", duration.String())

	if success {
		os.Exit(0)
	}

	os.Exit(1)
}
