package datastruct

import (
	"fmt"
	"reflect"
)
// Slices :Defines how to declare and use slices
func Slices() {
	// Slices are basically dynamic array implementation. In Go, we use slices more than arrays.
	// REMEMBER: slices points to same copy as array . So changes in slices are reflected in array
	fmt.Println(" \n***************** STARTING SLICES DEMO NOW  *****************")
	array1 := [6]int{2, 3, 4, 5, 67, 11}
	fmt.Println(array1, reflect.TypeOf(array1)) // [2 3 4 5 67 11] [6]int

	// How to create a slice ? ---- Mostly like an slice in python
	// Now for array1, we can create a slice like this:

	// METHOD1 : CREATE A SLICE FROM AN ARRAY
	s := array1[:]
	fmt.Println(s, reflect.TypeOf(s)) // []int

	s[2] = 25
	fmt.Print(array1, s, "\n") //[2 3 25 5 67 11] [2 3 25 5 67 11]

	// METHOD2: CREATE A SLICE DIRECTLY USING DIRECT DECLARATION
	/*  In array declaration, we have a specific size or  "..." as size argument inside [],
	but if square brackets is empty it's a slice */
	newSlice := []int{1, 2, 3, 4}
	fmt.Print(newSlice)

	// METHOD3: Use make function to create a slice
	var s2 []byte
	s2 = make([]byte, 5, 5) //len=5 and capacity=5

	fmt.Print(len(s2), cap(s2), "\n") // 5  5

	/* If we created a smaller slice earlier and want to increase it's size later,
	we can do it like s=s[:cap(s)] */

	// Gowing slices: Like dynamic array implementation in other languages.
	t := make([]int, len(s), (cap(s)+1)*2) //+1 in case cap is 0
	copy(t, s)
	s = t
	fmt.Print(s, len(s), cap(s)) // [2 3 25 5 67 11] 6 14

	//Appending data to end of a slice
	p := []string{"Hello", "Abhishek"}
	p = append(p, "Sharma")
	fmt.Print(p) // [Hello Abhishek Sharma]

fmt.Println("\n ***************** ENDING SLICES DEMO NOW  *****************")

}
