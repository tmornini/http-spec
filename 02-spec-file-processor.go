package main

import (
	"bufio"
	"net/http"
	"os"
)

func specFileProcessor(context context) {
	context.log("02 spec-file-processor")

	osFile, err := os.Open(context.Pathname)

	if errorHandler(&context, err) {
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

	context.WaitGroup.Done()
}
