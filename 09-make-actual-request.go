package main

import "net/http"

func makeActualRequest(context *context) bool {
	context.log("09 makeRequest")

	if context.HTTPClient == nil {
		context.HTTPClient = &http.Client{}
	}

	// func NewRequest(method, urlStr string, body io.Reader) (*Request, error)
	//     NewRequest returns a new Request given a method, URL, and optional body.
	//
	//     If the provided body is also an io.Closer, the returned Request.Body is set
	//     to body and will be closed by the Client methods Do, Post, and PostForm, and
	//     Transport.RoundTrip.
	//
	//     NewRequest returns a Request suitable for use with Client.Do or
	//     Transport.RoundTrip. To create a request for use with testing a Server
	//     Handler use either ReadRequest or manually update the Request fields. See
	//     the Request type's documentation for the difference between inbound and
	//     outbound request fields.

	request, err := http.NewRequest(
		context.Request.Verb,
		context.URIScheme+context.HostName+context.Request.Path,
		nil,
	)

	panicOn(err)

	response, err := context.HTTPClient.Do(request)

	panicOn(err)

	context.HTTPResponse = response

	return parseActualResponse(context)
}
