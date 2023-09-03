package main

import (
	"fmt"
	"strings"
)

func main() {
	var aString, aLowerString string

	fmt.Println("Enter a string")
	fmt.Scan(&aString)
	aLowerString = strings.ToLower(aString)

	if strings.HasPrefix(aLowerString, "i") && strings.HasSuffix(aLowerString, "n") && strings.Contains(aLowerString, "a") {
		fmt.Printf("Found!\n")
	} else {
		fmt.Printf("Not Found!\n")
	}
}
