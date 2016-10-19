package main

func processSpecFilesConcurrently(context *context) {
	context.log("01 process-spec-files-concurrently")

	for _, pathname := range context.Pathnames {
		context.Pathname = pathname

		context.WaitGroup.Add(1)

		// pass COPY of context (to goroutine)
		go processSpecFile(*context)
	}

	context.WaitGroup.Wait()

	// report per run stats

	exitWithSucess()
}
