package main

func processFile(context context) {
	context.log("03 processFile")

	defer context.WaitGroup.Done()

	context.Substitutions = map[string]string{}

	defer context.File.close()

	for {
		eof := parseRequestLine(&context)

		if eof {
			break
		}
	}
}
