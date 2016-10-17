package main

import "net/http"

func makeRequest(context *context) bool {
	context.log("9 makeRequest")

	if context.HTTPClient == nil {
		context.HTTPClient = &http.Client{}
	}

	request, err := http.NewRequest(
		context.Request.Verb,
		context.URIScheme+context.HostName+context.Request.Path,
		nil,
	)

	panicOn(err)

	response, err := context.HTTPClient.Do(request)

	panicOn(err)

	context.HTTPResponse = response

	defer response.Body.Close()

	// body, err := ioutil.ReadAll(response.Body)
	//
	// panicOn(err)

	return false
}
