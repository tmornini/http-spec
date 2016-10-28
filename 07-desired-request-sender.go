package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func desiredRequestMaker(context *context) {
	context.log("07-desired-request-maker")

	desiredRequest := context.SpecTriplet.DesiredRequest

	body := ioutil.NopCloser(strings.NewReader(desiredRequest.Body() + "\n"))

	request, err := http.NewRequest(
		desiredRequest.Method,
		context.URIScheme+context.HostName+desiredRequest.Path,
		body,
	)

	if errorHandler(context, err) {
		return
	}

	for _, headerLine := range context.SpecTriplet.DesiredRequest.HeaderLines {
		parts := strings.Split(headerLine.Text, ":")

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		request.Header.Add(key, value)
	}

	context.HTTPResponse, err = context.HTTPClient.Do(request)

	if errorHandler(context, err) {
		return
	}

	actualResponseReceiver(context)
}
