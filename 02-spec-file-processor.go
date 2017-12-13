package main

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"math/big"
	"net/http"
	"os"
)

func specFileProcessor(context context) {
	context.log("02 spec-file-processor")

	defer context.WaitGroup.Done()

	osFile, err := os.Open(context.Pathname)

	if errorHandler(&context, err) {
		context.ResultGathererChannel <- context

		return
	}

	space := new(big.Int).Exp(big.NewInt(62), big.NewInt(22), nil)
	uuid, err := rand.Int(rand.Reader, space)

	if err != nil {
		panic(err)
	}

	context.ID = uuid

	context.File = &file{
		bufio.NewReader(osFile),
		context.Pathname,
		osFile,
		0,
	}

	context.Substitutions = map[string]string{}
	context.Substitutions["YYYY-MM-DD"] = context.StartedAt.Format("2006-01-02")

	if context.SkipTLSVerification {
		context.HTTPClient = &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	} else {
		context.HTTPClient = &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	specTripletIterator(&context)

	osFile.Close()
}
