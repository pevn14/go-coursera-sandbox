package main

import (
	"encoding/json"
	"fmt"
	// "strings"
)

func main() {
	var aString string
	aMap := make(map[string]string)

	fmt.Println("Enter a name")
	fmt.Scan(&aString)
	aMap["Name"] = aString

	fmt.Println("Enter a address")
	fmt.Scan(&aString)
	aMap["Address"] = aString

	// fmt.Println("aMap:", aMap)

	b, err := json.Marshal(aMap)
	fmt.Println("Raw Json object: ", b)
	fmt.Println("Json object: ", string(b))
	fmt.Println("Marshal Error: ", err)
}
