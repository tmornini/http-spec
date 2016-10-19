package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func sendRequestReceiveResponse(context *context) error {
	context.log("07-send-request-receive-response")

	desiredRequest := context.SpecTriplet.DesiredRequest

	body := ioutil.NopCloser(strings.NewReader(desiredRequest.Body() + "\n"))

	request, err := http.NewRequest(
		desiredRequest.Method,
		context.URIScheme+context.HostName+desiredRequest.Path,
		body,
	)

	if err != nil {
		return err
	}

	for _, headerLine := range context.SpecTriplet.DesiredRequest.HeaderLines {
		parts := strings.Split(headerLine.Text, ":")

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		request.Header.Add(key, value)
	}

	// bytes, err := httputil.DumpRequestOut(request, true)
	//
	// if err != nil {
	// 	panic(err)
	// }
	//
	// fmt.Println(string(bytes))

	response, err := context.HTTPClient.Do(request)

	if err != nil {
		return err
	}

	context.HTTPResponse = response

	return translateResponse(context)
}
