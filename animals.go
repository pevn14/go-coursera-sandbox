package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
	
	// var aString string
	var cow, bird, snake Animal
	
	cow.Init("grass", "walk", "moo")         // init the animals 
	bird.Init("worms", "fly", "peep")
	snake.Init("mice", "slither", "hsss")
	
	fmt.Println("Enter the name of Animal (cow, bird or snake) and the request (eat, move or speak) separated by a space")
	for true {
		slice := prompt() //request for an animal 
		// fmt.Println("You entered:", slice[0], slice[1])
		switch slice[0] +"."+ slice[1] {
			case "cow.eat": fmt.Println(cow.Eat())
			case "cow.move": fmt.Println(cow.Move())
			case "cow.speak": fmt.Println(cow.Speak())
			case "bird.eat": fmt.Println(bird.Eat())
			case "bird.move": fmt.Println(bird.Move())
			case "bird.speak": fmt.Println(bird.Speak())
			case "snake.eat": fmt.Println(snake.Eat())
			case "snake.move": fmt.Println(snake.Move())
			case "snake.speak": fmt.Println(snake.Speak())
			default: fmt.Println("Bad request!")
		}
	}
}

// internal function, to prompt strings separated by a space
func prompt()[]string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)     // Remove the newline character at the end	
	return strings.Split(input, " ")     // Split the two strings using the space
}