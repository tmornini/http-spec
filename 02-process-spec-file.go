package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

// receive COPY of context (in goroutine)
func processSpecFile(context context) {
	context.log("02 process-spec-file")

	osFile, err := os.Open(context.Pathname)

	if err != nil {
		fmt.Println(context.File, "failure:", err)
	}

	defer osFile.Close()

	context.File = &file{
		bufio.NewReader(osFile),
		context.Pathname,
		osFile,
		0,
	}

	context.HTTPClient = &http.Client{}

	for {
		err = iterateSpecTriplets(&context)

		if err != nil {
			break
		}
	}

	switch err {
	case io.EOF:
		fmt.Println(context.File, "success")
	default:
		fmt.Println(context.File, "failure:", err)
	}

	context.WaitGroup.Done()
}
