package main

import (
	"fmt"
)

func main() {

	// initialise first
	var myArray [3]int
	// then declare
	myArray[0] = 12
	myArray[1] = 13
	myArray[2] = 14

	fmt.Println(myArray)
	fmt.Println(myArray[2])
}