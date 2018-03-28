package main

func iterateSpecFiles(s state) <-chan state {
	out := make(chan state)

	go func() {
		for _, pathname := range s.HTSFFilePaths {
			s.log("iterate-spec-files")

			s.HTSFFilePath = pathname
			out <- s
		}
		close(out)
	}()

	return out
}
