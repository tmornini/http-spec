package main

import (
	"fmt"
	"os"
)

func resultGatherer(context context) {
	// cannot log context here without overwriting context.Stage

	success := true

	successes := make(map[string]int)
	failures := make(map[string]int)

	for completedContext := range context.ResultGathererChannel {
		if completedContext.Err == nil {
			successes[completedContext.Stage]++
			fmt.Printf("[%s]: success!\n", context.Stage)
		} else {
			success = false
			failures[completedContext.Stage]++
			fmt.Printf("[%s]: failure: %s\n", completedContext.Stage, completedContext.Err.Error())
		}
	}

	fmt.Println("successes:", successes)
	fmt.Println("failures:", failures)

	if success {
		fmt.Println("http-spec: success!")
		os.Exit(0)
	} else {
		fmt.Println("http-spec: failure!")
		os.Exit(1)
	}
}
