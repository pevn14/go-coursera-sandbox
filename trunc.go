package main

import "fmt"

func main() {
    var aFloatingNumber float32

    fmt.Println("Enter a floating number")
    fmt.Scan(&aFloatingNumber)
    fmt.Printf("The scanned number is %d\n", int(aFloatingNumber))
}