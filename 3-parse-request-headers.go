package main

func parseRequestHeaders(context *context) {
	headerLines := []string{}

	for context.Scanner.Scan() {
		potentialHeaderLine := context.Scanner.Text()

		if potentialHeaderLine == "" {
			break
		}

		headerLines = append(headerLines, potentialHeaderLine)
	}

	context.Request.HeaderLines = &headerLines

	if err := context.Scanner.Err(); err != nil {
		panic(err)
	}

	parseRequestBody(context)
}
