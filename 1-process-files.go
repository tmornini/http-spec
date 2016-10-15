package main

func processFiles(https bool, hostname string, pathnames []string) {
	var prefix string

	if https {
		prefix = "https://" + hostname
	} else {
		prefix = "http://" + hostname
	}

	for _, pathname := range pathnames {
		context := &context{
			Pathname: pathname,
			Prefix:   prefix,
		}

		processFile(context)
	}
}
