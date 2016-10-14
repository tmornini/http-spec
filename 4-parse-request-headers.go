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

	if err := context.Scanner.Err(); err != nil {
		panic(err)
	}

	context.Request.HeaderLines = &headerLines

	parseRequestBody(context)
}
