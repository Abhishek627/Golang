// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(("hello world"))
	var a = 20
	fmt.Println(a)

	a = 30
	fmt.Println(a)

	//For loops

	b := 1
	for b <= 7 {
		fmt.Println(b)
		b = b + 1
	}

	for c := 10; c <= 12; c++ {
		fmt.Println(c)
	}

	//switch cases

	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	}

	//arrays

	var array [5]int
	fmt.Println(array)

	array[0] = 12
	fmt.Println(len(array))

	array1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array1)

	// Pointers

	var x int = 1
	var ip *int // ip is pointer to int
	ip = &x     // ip now points to x
	fmt.Print(ip)
	//new function is alternate way to create a variable. It returns the pointer to the variable instead of value.
	ptr := new(int)
	*ptr = 3

}
