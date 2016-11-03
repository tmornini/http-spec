package main

import (
	"fmt"
	"os"
	"time"
)

func resultGatherer(context context) {
	// cannot log context here without overwriting context.Stage

	success := true

	successes := 0
	failures := 0

	startedAt := time.Now()

	for completedContext := range context.ResultGathererChannel {
		duration := time.Since(completedContext.StartedAt)

		if completedContext.Err == nil {
			successes++
			fmt.Printf(
				"success %s %s\n",
				duration.String(),
				completedContext.SpecTriplet.String())
		} else {
			success = false
			failures++

			location := "failure " + duration.String() + " "

			if completedContext.File == nil {
				location += "[" + completedContext.Pathname + "]"
			} else {
				if completedContext.SpecTriplet == nil {
					location += completedContext.File.String()
				} else {
					location += completedContext.SpecTriplet.String()
				}
			}

			fmt.Println(location, completedContext.Err.Error())
		}
	}

	duration := time.Since(startedAt)

	fmt.Println("Total run time:", duration.String())

	if success {
		os.Exit(0)
	}

	os.Exit(1)
}
