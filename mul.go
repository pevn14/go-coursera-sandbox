package main

import (
	"fmt"
	"time"
	// "sync"
)

func prod(v1 int, v2 int, c chan int) {
	time.Sleep(1000 * time.Millisecond)    // pour verifier si le main attend bien la reception du message
	c <- v1 * v2
	fmt.Println("fin de la goroutine: ", v1, v2)
}

func main() {
	c := make(chan int)
	go prod(1, 2, c)
	go prod(3, 4, c)
	fmt.Println("waiting messages")
	a := <- c
	fmt.Println("receive from first: ", a)
	b := <- c
	fmt.Println("receive from second: ", b)
	fmt.Println(a*b)
}
