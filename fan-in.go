package main

import "sync"

func fanIn(in ...<-chan state) <-chan state {
	var wg sync.WaitGroup
	out := make(chan state)

	output := func(in <-chan state) {
		for s := range in {
			out <- s
		}
		wg.Done()
	}

	wg.Add(len(in))

	for _, s := range in {
		go output(s)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
