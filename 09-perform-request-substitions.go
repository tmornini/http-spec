package main

func performRequestSubstitutions(context *context) bool {
	context.log("09 performRequestSubstitutions")

	for _, line := range context.Request.Lines {
		line.substitute(context)
	}

	return performExpectedResponseSubstitutions(context)
}
