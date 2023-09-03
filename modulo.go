package main

import (
	"fmt"
)

type Philo struct {
	id int
	c, d chan int
}

func main(){
	for i := 0; i < 3*5; i= i+1  {
		fmt.Println("i: ", i)
		fmt.Println("i %5: ", i % 5)
		fmt.Println("i+2 %5: ", (i+2) % 5)
		fmt.Println("")
	}
}