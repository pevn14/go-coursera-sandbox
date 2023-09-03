package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a, v0, s0 float64) func(float64)float64 {
	return func (t float64)float64 {							// use anonymous function
		return 0.5 * a *t * t + v0 * t + s0
	}
}

func main() {
	// first step, enter the parameters
	var aString string
	fmt.Println("Enter an acceleraion")
	fmt.Scan(&aString)
	a, _ := strconv.ParseFloat(aString, 64)
	fmt.Println("Enter an initial velocity")
	fmt.Scan(&aString)
	v, _ := strconv.ParseFloat(aString, 64)
	fmt.Println("Enter an initial displacement")
	fmt.Scan(&aString)
	s, _ := strconv.ParseFloat(aString, 64)
	
	// a, v, s := 10., 2., 1.   // only for testing the call

	// then generate the function
	fn :=  GenDisplaceFn(a, v, s)

	// enter the time to compute fn(time)
	fmt.Println("Enter a time")
	fmt.Scan(&aString)
	t, _ := strconv.ParseFloat(aString, 64)

	// and print the result to call f()
	fmt.Printf("Parametrs: a= %.2f, v0= %.2f, s0= %.2f\n", a, v, s)     
	//fmt.Printf("Result for t= %.2f : %.2f\n", t, fn(10))  // only for testing
	fmt.Printf("Result for t= %.2f : %.2f\n", t, fn(t))
}