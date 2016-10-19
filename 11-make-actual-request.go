package main

import "net/http"

func makeActualRequest(context *context) bool {
	context.log("11 makeActualRequest")

	if context.HTTPClient == nil {
		context.HTTPClient = &http.Client{}
	}

	request, err := http.NewRequest(
		context.Request.Verb,
		context.URIScheme+context.HostName+context.Request.Path,
		nil,
	)

	exitWithStatusOneIf(err)

	response, err := context.HTTPClient.Do(request)

	exitWithStatusOneIf(err)

	context.HTTPResponse = response

	return parseActualStatusLine(context)
}
