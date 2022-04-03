package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done chan interface{}, pulseInterval time.Duration) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{})
		results := make(chan time.Time)

		go func() {
			defer close(heartbeat)
			defer close(results)

			// time.Tick uses NewTicker
			// NewTicker returns a new Ticker containing a channel that will
			// sendthe time on the channel after each tick
			pulse := time.Tick(pulseInterval)
			workGen := time.Tick(2 * pulseInterval)

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				// Default case needs to be added because no one may be
				// listening to our heartbeat. The results emitted from the
				// goroutine are critical but the pulses are not.
				default:
				}
			}

			sendResult := func(r time.Time) {
				// since we might require sending out multiple pulse while
				// we wait for result, we need to wrap it in for loop.
				for {
					select {
					case <-pulse:
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse:
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()

		return heartbeat, results
	}

	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() {
		close(done)
	})

	const timeout = 2 * time.Second
	heartbeat, results := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("result %v\n", r.Second())
		case <-time.After(timeout):
			return
		}
	}
}
