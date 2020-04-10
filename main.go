package main

import (
	"fmt"
	"reflect"
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

	// Here's another syntax
	my2ndArray := [3]int{18, 25, 30}
	fmt.Println(my2ndArray)

	// here the "..." tells golang to work out the array length for
	// us based on the number or values provided.
	// This technique is referred to as array literals
	stooges := [...]string{"Moe", "Larry", "Curly", "charlie"}
	fmt.Println(reflect.TypeOf(stooges)) // [4]string
	fmt.Println(len(stooges))            // 4

}

// https://yourbasic.org/golang/three-dots-ellipsis/
