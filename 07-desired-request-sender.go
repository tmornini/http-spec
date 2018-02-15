package main

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func desiredRequestSender(context *context) {
	context.log("07-desired-request-sender")

	desiredRequest := context.SpecTriplet.DesiredRequest

	body := ioutil.NopCloser(strings.NewReader(desiredRequest.Body()))

	request, err := http.NewRequest(
		desiredRequest.Method(),
		desiredRequest.AbsoluteURI(),
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

	attempt := 0

	for {
		context.HTTPResponse, err = context.HTTPClient.Do(request)

		attempt++

		if err == nil {
			break
		}

		if attempt >= context.MaxHTTPAttempts {
			context.Err = err

			return
		}

		time.Sleep(context.HTTPRetryDelay)

		continue
	}

	actualResponseReceiver(context)
}
