package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }
func (c Cow) Eat()   { println("Grass") }
func (c Cow) Move()  { println("Walk") }
func (c Cow) Speak() { println("Moo") }

type Bird struct{ name string }
func (c Bird) Eat()   { println("Worms") }
func (c Bird) Move()  { println("fly") }
func (c Bird) Speak() { println("peep") }

type Snake struct{ name string }
func (c Snake) Eat()   { println("mice") }
func (c Snake) Move()  { println("slither") }
func (c Snake) Speak() { println("hss") }

func main() {

	myAnimalsList := make(map[string]Animal)

	myAnimalsList["Rose"] = Cow{"Rose"}
	myAnimalsList["Lotta"] = Bird{"Lotta"}
	myAnimalsList["Vegas"] = Snake{"Vegas"}

	fmt.Printf("myAnimal: %+v\n", myAnimalsList["Rose"])
	fmt.Printf("myAnimal Type: %T\n", myAnimalsList["Rose"])
	fmt.Printf("myAnimal: %+v\n", myAnimalsList["Lotta"])
	fmt.Printf("myAnimal Type: %T\n", myAnimalsList["Lotta"])
	fmt.Printf("myAnimal Name: %s\n", myAnimalsList["Rose"])
	myAnimalsList["Rose"].Eat()
	myAnimalsList["Lotta"].Eat()
	
	myAnimal := "Vegas"
	myAnimalsList[myAnimal].Eat()
	
}
