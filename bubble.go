package main

import (
	"fmt"
	"strings"
	"strconv"
	"reflect"
)

func BubbleSort(sli []int) {
	for i:=0; i<len(sli)-1; i++ {
		if sli[i] > sli[i+1] {
			Swap(sli, i)
			BubbleSort(sli)
		}			
	}
}

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func ConvertSliceStringToInt(strings []string) []int {
	var ints []int
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func main() {
	var aString string
	fmt.Println("Enter 10 integers separated by comma without space")
	fmt.Scanln(&aString)
	// aString = "14,2,7,1,15,45,5,-1,125,0"  // for testing

	// first extract a slice of int from the string
	parts := strings.Split(aString, ",")
	ints := ConvertSliceStringToInt(parts)
	fmt.Printf("Check the type of ints: %v\n", reflect.TypeOf(ints).Kind())  // check the type of ints ; it must be a slice and not an array
	
	// process the bubble sorting
	fmt.Println("Before sorting: ", ints)
	BubbleSort(ints)
	fmt.Println("After sorting: ", ints)
}

