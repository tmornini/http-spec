package main

func specFileScatter(context *context) {
	context.log("01 spec-file-scatter")

	for _, pathname := range context.Pathnames {
		context.Pathname = pathname

		context.WaitGroup.Add(1)
		go specFileProcessor(*context)
	}
}
