package main

import (
	"fmt"
	// "math/rand"
	"strings"
	"strconv"
)

func main() {
	//Generate a random array of length n for debug
	// size := 18     // size of the initial debug array
	// sli := rand.Perm(size)   // Perm returns, as a slice of n ints, a pseudo-random permutation of the integers in the half-open interval [0,n) from the default Source.
	
	var aString string
	fmt.Println("Enter a list of integers separated by comma without space")
	fmt.Scanln(&aString)
	// first extract a slice of int from the string
	parts := strings.Split(aString, ",")
	sli := convertSliceStringToInt(parts)
	
	
	fmt.Println("the initial list: ", sli)

	n := 4         					// number of subroutines
	sizeSub :=  int(len(sli) / n)		// size of subslice
   

	c := make(chan []int)        // create a channel and run the goroutines
		for i:=0; i<n; i++ {
		go sort(sli[i*sizeSub:(i+1)*sizeSub], i, c) 
	}
	
	// wait the messages from each goroutine
	var ints []int
	for i:=0; i<n; i++ {
		mes := <- c
		ints = append(ints, mes...) 
	}
	fmt.Println("The final merged list: ", ints)  //note: this list is simply an merging of the 4 sub-lists
	bubbleSort(ints)
	fmt.Println("The final sorted list: ", ints) 
}

func sort(sli []int, id int, c chan []int) {
	// fmt.Println("sub slice not sorted: ", id, sli)    
	bubbleSort(sli)
	fmt.Printf("sub slice sorted by goroutine #%d : %d\n", id, sli)
	c <- sli    
}

func bubbleSort(sli []int) {             // reuse this function from an old exercice
	for i:=0; i<len(sli)-1; i++ {
		if sli[i] > sli[i+1] {
			sli[i], sli[i+1] = sli[i+1], sli[i]			
			bubbleSort(sli)
		}			
	}
}

func convertSliceStringToInt(strings []string) []int {
	var ints []int
	for _, s := range strings {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}