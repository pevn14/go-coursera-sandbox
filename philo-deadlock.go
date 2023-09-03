package main

import (
	"fmt"
	"sync"
)

type ChopS struct{ sync.Mutex }

type Philo struct {
	leftCS, rightCS *ChopS
}

func (p Philo) eat(wg *sync.WaitGroup) {
	for {
		p.leftCS.Lock() // prend les 2 baguettes
		p.rightCS.Lock()

		fmt.Println("eating") // mange

		p.rightCS.Unlock() // pose les baguettes
		p.leftCS.Unlock()
		// wg.Done()   // l'attente du main() sera infinie pour provoquer le deadlock
		break
	}
}

func main() {
	fmt.Println("Creation des baguettes")
	CSticks := make([]*ChopS, 5) //creation des baguettes
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	fmt.Println("Creation des philosophes")
	philos := make([]*Philo, 5) //creation des philosophes

	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5]}
	}

	fmt.Println(philos)

	fmt.Println("Lancement des philosophes")
	var wg sync.WaitGroup // create a waiting group
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philos[i].eat(&wg)
	}
	fmt.Println("Attente des done")
	wg.Wait() //waiting
	fmt.Println("Fin de l'attente")
}
