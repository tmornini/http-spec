package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func desiredRequestSender(context *context) {
	context.log("07-desired-request-sender")

	desiredRequest := context.SpecTriplet.DesiredRequest

	body := ioutil.NopCloser(strings.NewReader(desiredRequest.Body() + "\n"))

	fmt.Println(desiredRequest.Method(), context.URLPrefix+desiredRequest.Path())

	request, err := http.NewRequest(
		desiredRequest.Method(),
		context.URLPrefix+desiredRequest.Path(),
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

	context.SpecTriplet.StartedAt = time.Now()

	context.HTTPResponse, err = context.HTTPClient.Do(request)

	if errorHandler(context, err) {
		return
	}

	actualResponseReceiver(context)
}
