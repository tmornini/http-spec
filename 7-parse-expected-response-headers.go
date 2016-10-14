package main

func parseExpectedResponseHeaders(context *context) {
	headerLines := []*headerLine{}

	for context.Scanner.Scan() {
		potentialHeaderLineInput := context.Scanner.Text()

		if potentialHeaderLineInput == "" {
			break
		}

		headerLine := &headerLine{potentialHeaderLineInput}

		headerLines = append(headerLines, headerLine)
	}

	err := context.Scanner.Err()

	if err != nil {
		panic(err)
	}

	context.ExpectedResponse.HeaderLines = &headerLines

	parseExpectedResponseHeaderRegexps(context)
}
