package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func sendRequest(state *state) error {
	state.log("send-request")

	desiredRequest := state.SpecTriplet.DesiredRequest

	body := ioutil.NopCloser(strings.NewReader(desiredRequest.Body()))

	request, err := http.NewRequest(
		desiredRequest.Method(),
		desiredRequest.AbsoluteURI(),
		body,
	)

	if err != nil {
		return err
	}

	for _, headerLine := range state.SpecTriplet.DesiredRequest.HeaderLines {
		parts := strings.Split(headerLine.Text, ":")

		if len(parts) < 2 {
			return fmt.Errorf("invalid header line: %s", headerLine.String())
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(strings.Join(parts[1:], ""))

		request.Header.Add(key, value)
	}

	state.SpecTriplet.StartedAt = time.Now()

	attempt := 0

	for {
		state.HTTPResponse, err = state.HTTPClient.Do(request)

		attempt++

		if err == nil {
			break
		}

		if attempt >= state.MaxHTTPAttempts {
			return err
		}

		time.Sleep(state.HTTPRetryDelay)

		continue
	}

	return parseActualResponse(state)
}
