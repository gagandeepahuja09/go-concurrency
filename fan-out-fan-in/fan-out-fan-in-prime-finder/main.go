package main

import (
	"runtime"
	"sync"
)

func main() {
	// FanOut Logic
	numFinder := runtime.NumCPU()
	finders := make([]<-chan int, numFinder)
	for i := 0; i < numFinder; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}

	// FanIn Logic
	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})

		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		for _, c := range channels {
			wg.Add(1)
			go multiplex(c)
		}

		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}
}
