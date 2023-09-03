package main

import (
	"fmt"
)

type Animal struct {
	food string
	locomotion string
	spoken string
}

func (a *Animal) Init(f, l, s string) {
	a.food = f
	a.locomotion = l
	a.spoken = s
}

func (a Animal) Eat()string {
	return a.food
}

func (a Animal) Move()string {
	return a.locomotion
}

func (a Animal) Speak()string {
	return a.spoken
}

func main() {
	
	var animal, action string
	var cow, bird, snake Animal

	cow.Init("grass", "walk", "moo") 		// initialisation of the animals
	bird.Init("worms", "fly", "peep")
	snake.Init("mice", "slither", "hsss")

	myAnimals := map[string]Animal{
		"cow":   cow,
		"bird":  bird,
		"snake": snake,
	}
	
	fmt.Println("Enter the name of Animal (cow, bird or snake) and the request (eat, move or speak) separated by a space")
	for true {
		fmt.Scanf("%s %s", &animal, &action)
		fmt.Println("You entered:", animal, action)
			switch action {
				case "eat": fmt.Println(myAnimals[animal].Eat())
				case "move": fmt.Println(myAnimals[animal].Move())
				case "speak": fmt.Println(myAnimals[animal].Speak())
				default: fmt.Println("Bad request!")
			}
	}
}