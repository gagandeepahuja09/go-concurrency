package main

func main() {
	// repeatFn will keep on repeatedly calling the fn function and write the value
	// returned to valueStream.
	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
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

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	// primeFinder := func(done <-chan interface{}, intStream <-chan int) <-chan interface{} {
	// 	primeStream := make(chan interface{})
	// 	go func() {
	// 		defer close(primeStream)

	// 		for integer := range intStream {

	// 		}
	// 	}()
	// 	return primeStream
	// }
}
