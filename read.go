package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	// Prompt the user for the file name
	fmt.Print("Please enter the name of the file: ")
	var fileName string
	fmt.Scanln(&fileName)

	// Open the file in read mode
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening the file: %v", err)
	}
	defer file.Close()

	var people []Person

	// Use a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) == 2 {
			person := Person{
				firstName: parts[0],
				lastName:  parts[1],
			}
			people = append(people, person)
		} else {
			fmt.Printf("The line '%s' does not have exactly two strings separated by a space.\n", line)
		}
	}

	// Handle potential errors from the scanner
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading the file: %v", err)
	}

	// Display the information for each person
	for _, p := range people {
		fmt.Printf("First Name: %s, Last Name: %s\n", p.firstName, p.lastName)
	}
}
