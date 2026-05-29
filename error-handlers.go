package main

func errorHandler(context *context, err error) bool {
	if err != nil {
		context.Err = err

		return true
	}

	return false
}
