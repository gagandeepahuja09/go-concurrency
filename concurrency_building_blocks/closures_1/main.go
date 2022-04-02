package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		// use the same address space rather than a copy.
		salutation = "welcome"
	}()
	wg.Wait() // this is the join point.
	fmt.Println("Salutation: ", salutation)
}
