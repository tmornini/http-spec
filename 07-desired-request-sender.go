package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func desiredRequestSender(context *context) {
	context.log("07 desired-request-sender")

	desiredRequest := context.SpecTriplet.DesiredRequest

	body := io.NopCloser(strings.NewReader(desiredRequest.Body()))

	request, err := http.NewRequest(
		desiredRequest.Method(),
		desiredRequest.URL(),
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

		if err == nil && !context.isRetryStatusCode() {
			break
		}

		if attempt >= context.MaxHTTPAttempts {
			if err != nil {
				context.Err = err
			} else {
				context.Err = fmt.Errorf(
					"retry limit reached on status %d",
					context.HTTPResponse.StatusCode,
				)
			}

			return
		}

		if err == nil {
			context.HTTPResponse.Body.Close()
		}

		time.Sleep(context.HTTPRetryDelay)
	}

	actualResponseTranslator(context)
}
