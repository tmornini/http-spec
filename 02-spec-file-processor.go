package main

import (
	"bufio"
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
	context.File = &file{
		bufio.NewReader(osFile),
		context.Pathname,
		osFile,
		0,
	}

	context.Substitutions = map[string]string{}

	context.HTTPClient = &http.Client{}

	specTripletIterator(&context)

	osFile.Close()
}
