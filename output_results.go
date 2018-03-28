package main

import (
	"fmt"
	"time"
)

func outputResults(in <-chan state) <-chan int {
	out := make(chan int)

	go func() {
		successCount := 0
		failureCount := 0
		startedAt := time.Now()

		for s := range in {
			s.log("output-results")

			if s.Result.Failure {
				failureCount++
			} else {
				successCount++
			}

			fmt.Print(s.Result.Message)
		}

		duration := time.Since(startedAt)

		// fmt.Println()

		if failureCount > 0 {
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

			out <- 1

		} else {
			fmt.Printf(
				"\n%sSUCCESS:%s %s+%d%s %s\n",
				Green,
				Reset,
				Green,
				successCount,
				Reset,
				duration.String(),
			)
		}

		// os.Exit(0)
		out <- 0

		close(out)
	}()

	return out
}
