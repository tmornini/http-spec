package main

import (
	"bufio"
	"crypto/rand"
	"crypto/tls"
	"math/big"
	"net/http"
	"os"
)

const idBase = 62
const idLength = 22

func specFileProcessor(context context) {
	context.log("02 spec-file-processor")

	defer context.WaitGroup.Done()

	osFile, err := os.Open(context.Pathname)

	if errorHandler(&context, err) {
		context.ResultGathererChannel <- context

		return
	}

	space := new(big.Int).Exp(
		big.NewInt(idBase),
		big.NewInt(idLength),
		nil,
	)
	uuid, err := rand.Int(rand.Reader, space)

	if errorHandler(&context, err) {
		context.ResultGathererChannel <- context

		return
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

	randomID, err := newRandom(&context)

	if errorHandler(&context, err) {
		context.ResultGathererChannel <- context

		return
	}

	context.Substitutions["random-uuid"] = randomID

	if context.SkipTLSVerification {
		context.HTTPClient = &http.Client{
			Timeout: context.HTTPTimeout,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	} else {
		context.HTTPClient = &http.Client{
			Timeout: context.HTTPTimeout,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	specTripletIterator(&context)

	osFile.Close()
}
