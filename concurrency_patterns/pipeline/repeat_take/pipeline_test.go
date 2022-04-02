package main

import "testing"

// 2256 ns/op
func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()

	for range toString(done, take(done, repeat(done, "gagan,", "gagandeep,"), b.N)) {
	}
}

// 389.6 ns/op
func BenchmarkTyped(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	repeatString := func(done <-chan interface{}, values ...string) <-chan string {
		valueStream := make(chan string)
		go func() {
			defer close(valueStream)
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}()
		return valueStream
	}

	takeString := func(done <-chan interface{},
		valueStream <-chan string, num int) <-chan string {
		takeStream := make(chan string)
		go func() {
			// don't forget to close it.
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	b.ResetTimer()

	for range takeString(done, repeatString(done, "gagan,", "gagandeep,"), b.N) {
	}
}
