package main

import (
	"fmt"
)

type Animal interface {						// define the interface for Cow, Bird an Snake
	Eat()
	Move()
	Speak()
}

type Cow struct {name string}              // define the Cow object
func (c Cow) Eat() {println("Grass")} 
func (c Cow) Move() {println("Walk")} 
func (c Cow) Speak() {println("Moo")} 

type Bird struct {name string}            // define the Bird object
func (c Bird) Eat() {println("Worms")} 
func (c Bird) Move() {println("fly")} 
func (c Bird) Speak() {println("peep")}

type Snake struct{name string}           // define the Snake object
func (c Snake) Eat() {println("mice")}
func (c Snake) Move() {println("slither")}
func (c Snake) Speak() {println("hsss")}

func main() {
	
	var command, name, parameter string
	// var err error
	
	myAnimalsList := make(map[string]Animal)   // define a map of Animal
    	
	fmt.Println("")
	fmt.Println("Create a new animal : type the command (newanimal), the animal name, and its type (cow, bird or snake) separated by a space")
	fmt.Println("Or create a request : type the command (request), the animal name, and the request ((eat, move or speak) separated by a space")
	
	for true {
		fmt.Printf("> ")
		_, err := fmt.Scanf("%s %s %s\n", &command, &name, &parameter)
		if (err != nil) {
			fmt.Println("Please enter only 3 strings separated by a space!")
			} else {
			fmt.Println("You entered:", command, name, parameter )
			switch command {
				case "newanimal":
					fmt.Println("Create a new animal:", name, parameter)
					switch parameter{
						case "cow": myAnimalsList[name] = Cow{name} 
						case "bird": myAnimalsList[name] = Bird{name}
						case "snake": myAnimalsList[name] = Snake{name}
						default: fmt.Println("Bad animal: cow, bird or snake!")	
					}
				case "request":
					animal, exists := myAnimalsList[name]
    				if !exists {
        				fmt.Println(name,": does not exist!")
        				break
    				}
    				fmt.Println("Request an animal:", name, parameter)
    				switch parameter {
    					case "eat": animal.Eat()
    					case "move": animal.Move()
    					case "speak": animal.Speak()
    				default:
        				fmt.Println("Bad request: only eat, move or speak!")	
    				}
				default:
					fmt.Println("Bad command!")		
			}	
		}
	}
}
