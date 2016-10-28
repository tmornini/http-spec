package main

import "io"

func errorHandler(context *context, err error) bool {
	if err != nil {
		context.Err = err

		return true
	}

	return false
}

func errorNotEOFHandler(context *context, err error) bool {
	if err != nil && err != io.EOF {
		return errorHandler(context, err)
	}

	return false
}
