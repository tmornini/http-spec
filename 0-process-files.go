package main

func processFiles(https bool, hostname string, pathnames []string) {
	var prefix string

	switch https {
	case true:
		prefix = "https://" + hostname
	case false:
		prefix = "http://" + hostname
	}

	for _, pathname := range pathnames {
		context := &context{
			Pathname: pathname,
			Request: &request{
				Prefix: prefix,
			},
		}
		processFile(context)
	}
}
