package main

import "fmt"

func main() {
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}

	fmt.Println(counter)
}

/*
A race condition occurs when multiple threads (or goroutines in Go) access shared data and try to change it at the same time. In this case, multiple goroutines are trying to increment the counter variable concurrently without any synchronization mechanism.
*/
