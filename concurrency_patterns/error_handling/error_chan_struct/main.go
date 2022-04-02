package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) chan Result {
		responses := make(chan Result)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case responses <- result:
				}
			}
		}()
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	for result := range checkStatus(done, "https://www.google.com", "https://badhost") {
		if result.Error != nil {
			fmt.Printf("Error: %v\n", result.Error)
		} else {
			fmt.Printf("Response: %v\n", result.Response.StatusCode)
		}
	}
}
