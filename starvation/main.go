package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var sharedLock sync.Mutex

const runtime = 1 * time.Second

func greedyWorker() {
	defer wg.Done()
	count := 0
	for begin := time.Now(); time.Since(begin) < runtime; {
		sharedLock.Lock()
		time.Sleep(3 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}
	fmt.Printf("Greedy worker was able to execute %v worker loops\n", count)
}

func politeWorker() {
	defer wg.Done()
	count := 0
	for begin := time.Now(); time.Since(begin) < runtime; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}
	fmt.Printf("Polite worker was able to execute %v worker loops\n", count)
}

func main() {
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
