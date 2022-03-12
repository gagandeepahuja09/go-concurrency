package main

import "fmt"

func main() {
	i := 0
	go func() {
		i++
	}()
	fmt.Printf("value of i is : %d\n", i)
}
