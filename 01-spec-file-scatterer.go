package main

func specFileScatterer(context *context) {
	context.enterStage("01 spec-file-scatterer")

	for _, pathname := range context.Pathnames {
		context.Pathname = pathname

		context.WaitGroup.Add(1)
		go specFileProcessor(*context)
	}
}
