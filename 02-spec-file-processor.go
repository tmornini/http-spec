package main

import (
	"bufio"
	"os"
)

const idBase = 62
const idLength = 22

func specFileProcessor(context context) {
	context.enterStage("02 spec-file-processor")

	defer context.WaitGroup.Done()

	osFile, err := os.Open(context.Pathname)

	if errorHandler(&context, err) {
		context.ResultGathererChannel <- context

		return
	}

	randomID, err := newRandom(&context)

	if errorHandler(&context, err) {
		context.ResultGathererChannel <- context

		return
	}

	context.ID = randomID

	context.File = &file{
		bufio.NewReader(osFile),
		context.Pathname,
		osFile,
		0,
	}

	context.Substitutions = map[string]string{}
	context.Substitutions["YYYY-MM-DD"] = context.StartedAt.Format("2006-01-02")
	context.Substitutions["random-uuid"] = randomID

	context.HTTPClient = newHTTPClient(
		context.SkipTLSVerification,
		context.HTTPTimeout,
	)

	specTripletIterator(&context)

	osFile.Close()
}
