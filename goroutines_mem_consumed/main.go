package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()

		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup

	noop := func() { wg.Done(); <-c }

	const numGoRoutines = 1e4
	wg.Add(numGoRoutines)
	before := memConsumed()
	for i := numGoRoutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	// took only .211kb.
	fmt.Printf("%.3fkb", float64(after-before)/numGoRoutines/1000)
}
