package main

import (
	"bufio"
	"os"
)

func processSpecFile(in <-chan state) <-chan state {
	out := make(chan state)

	go func() {
		for s := range in {
			s.log("process-spec-file")

			// Handle any errors
			if s.Error != nil {
				out <- s
				continue
			}

			id, err := newRandom()

			if err != nil {
				s.Error = err
				out <- s

				continue
			}

			s.ID = id

			osFile, err := os.Open(s.HTSFFilePath)

			if err != nil {
				s.Error = err
				out <- s

				continue
			}

			s.File = &file{
				bufio.NewReader(osFile),
				s.HTSFFilePath,
				osFile,
				0,
			}

			out <- s
		}

		close(out)
	}()

	return out
}
