package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

func newHTTPClient(skipTLS bool, timeout time.Duration) *http.Client {
	client := &http.Client{
		Timeout: timeout,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	if skipTLS {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}

	return client
}
