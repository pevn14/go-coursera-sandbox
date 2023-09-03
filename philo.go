package main

import (
	"fmt"
	"sync"
)

// mutexes are no longer used, as the host controls the scheduling

type Philo struct {
	id int          // just an id for the print
	nb int			// counter of meal
	c, d chan int   // 2 channels to sync with host
	mu sync.Mutex 
}

func (p *Philo) eat() {
	for {																// mutexes are no longer used, as the host controls the scheduling
		<- p.c  														// waiting a sync message from the host
		fmt.Println("Starting eating ", p.id)
		p.d <- 0  														// send a sync message to the host
	}																	// the host controls the end of the main() => no end of loop and no wg.Done() here
}

func host(p []*Philo, wg *sync.WaitGroup, maxEat int) {
	
	fmt.Println("")
	fmt.Println("Start dining for all philosophers")
	fmt.Println("")
	
	for i := 0; i <= (maxEat -1) * 5 ; i++  {    								// 10 = (nb of meal -1) * nb of philosopers
		
		// the tests are compulsory, because on the last lap of the loop there 
		// may be just one philosopher who hasn't eaten the expected number of times		
		philosophe := p[i%5]
		voisin := p[(i+2)%5]

		philosophe.mu.Lock()
		if (philosophe.nb) <= maxEat {
			philosophe.c <- 0        											// send a sync message to the first philosopher
			<- philosophe.d          											// waiting for the first philosopher 
			fmt.Println("Finishing eating ", philosophe.id, " ; Number is ", philosophe.nb)
			philosophe.nb ++
		}
		philosophe.mu.Unlock()
		
		voisin.mu.Lock()
		if (voisin.nb) <= maxEat {
			voisin.c <- 0
			<-voisin.d      											// waiting for the second philosopher  
			fmt.Println("Finishing eating ", voisin.id, " ; Number is ", voisin.nb)
			voisin.nb ++    											// an another to the second plilosopher
		}
		voisin.mu.Unlock()

		// fmt.Println("") 
	}
	fmt.Println("End dining for all philosophers")
	wg.Done()
}

func main() {

	philos := make([]*Philo, 5)   					//creation of the philosophers

	for i := 0; i < 5; i++ {
		c := make(chan int)		
		d := make(chan int)		
		philos[i] = &Philo{i+1, 1, c, d, sync.Mutex{}}    		   // id of the first philosopher is 1, not 0
	}
	
	fmt.Println("Lauching the philosophers")
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	
	fmt.Println("Launching the host")
	var wg sync.WaitGroup                		// create a waiting group only for the host
	wg.Add(1)
	go host(philos, &wg, 3)  // last parameter is the number of east, we can try with anothe value if its work

	fmt.Println("Waiting for the done by host")
	wg.Wait()									//waiting
	fmt.Println("End of main()")
}
