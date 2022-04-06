package test_end_of_work

import (
	"testing"
	"time"
)

func DoWork(done <-chan interface{}, nums ...int) (<-chan interface{}, <-chan int) {
	// buffer of 1 ensures that there is always a pulse of 1 sent out
	// irrespective of anyone listening on time.
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second)

		for _, n := range nums {
			// seperate select block for heartbeat because this should happen
			// irrespective of whether the result is sent.
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case intStream <- n:
			}
		}
	}()

	return heartbeat, intStream
}

// this is a bad test as it's non-deterministic. It's assuming that it would finish in
// 1 second. We can increase the timeout but that would slow down the test suite.
func TestDoWork_GeneratesAllNumbers_ND(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for _, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf("test failed: %v != %v", r, expected)
			}
		case <-time.After(1 * time.Second):
			t.Fatal("test timed out")
		}
	}
}

func TestDoWork_GeneratesAllNumbers_Deterministic(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := DoWork(done, intSlice...)

	<-heartbeat

	i := 0
	for r := range results {
		if r != intSlice[i] {
			t.Errorf("test failed: %v != %v", r, intSlice[i])
		}
		i++
	}
}
