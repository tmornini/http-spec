package main

func processFiles(https bool, hostname string, pathnames []string) {
	var uriScheme string

	if https {
		uriScheme = "https://" + hostname
	} else {
		uriScheme = "http://" + hostname
	}

	for _, pathname := range pathnames {
		context := &context{
			Pathname:  pathname,
			UriScheme: uriScheme,
		}

		processFile(context)
	}
}
