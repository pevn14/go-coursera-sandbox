package main

import (
	"fmt"
	"time"
	"sync"
)

var x int

func add(wg *sync.WaitGroup, id string) {
	for true {
		var mut sync.Mutex
		mut.Lock()           // mutual exclusion area
		x += 1
		mut.Unlock()

		// fmt.Printf("I'm %s, and x= %d\n", id, x)   //u ncomment to see that the order of execution of the go routines is totally random
		if x == 1000 {
			fmt.Printf("I'm %s, and I WIN with x= %d\n", id, x)
			wg.Done()
			break
		}
		if x >= 1000 {
			fmt.Printf("I'm %s, and I Lost with x= %d\n", id, x)
			wg.Done()
			break
		}
	 	time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	x = 0
	// each goroutine increments the global counter x
	// the winner is the one who increases x to 100
	// and the other break
	// totally useless program, but illustrates the race condition
	// which process well be the Winner ?
	
	var wg sync.WaitGroup                // create a waiting group
	
	wg.Add(1); go add(&wg,"a")
	
	wg.Add(1)
	go add(&wg,"b")

	wg.Add(1)
	go add(&wg,"c")

	wg.Add(1)
	go add(&wg,"d")

	wg.Add(1)
	go add(&wg,"e")

	wg.Add(1)
	go add(&wg,"f")

	fmt.Println("Waiting")
	wg.Wait()					//waiting
	fmt.Printf("x = %d at the end\n", x) // check the value a the end

}
