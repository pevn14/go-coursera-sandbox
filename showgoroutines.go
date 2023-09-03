package main

import (
    "fmt"
    "time"
)

var x int

func add(id string) {
	for true {
		time.Sleep(10 * time.Millisecond)
		x = x + 1
		// fmt.Printf("I'm %s, and x= %d\n", id, x)   //u ncomment to see that the order of execution of the go routines is totally random
		if x == 100 {
			fmt.Printf("I'm %s, and I WIN with x= %d\n", id,x)
			break
		}
		if x >= 100 {
			fmt.Printf("I'm %s, and I Lost with x= %d\n", id,x)
			break
		}
	}	
}

func main() {
	x = 0
	// each goroutine increments the global counter x
	// the winner is the one who increases x to 100
	// and the other break
	// totally useless program, but illustrates the race condition
	// which process well be the Winner ?
	go add("a")
    go add("b")
    go add("c")
    go add("d")
    go add("e")
    go add("f")
	time.Sleep(1 * time.Second)  // otherwise program stops immediately 
	fmt.Printf("x = %d at the end\n", x)   // check the value a the end

}