package main	

import (
	"fmt"
	"math"
	"strings"
)

type circle struct {
	radius float32
}

type rectangle struct {
	length , breath float32
}

// This is a method with receiver type circle struct. This receives a copy of the struct, can't modify the struct directy
func (c circle) Area() float32 {
	return math.Pi * c.radius * c.radius
}

// We can have multiple methods with same name as long as receiver type is different.
func (r rectangle) Area() float32 {
	return r.length * r.breath
}
/***************************************************************************************/

// To modify the value of struct we can pass a pointer to struct as type
func (c *circle) changeRadius(newrad float32) {
	c.radius= newrad
}

/***************************************************************************************/

type mystring string

func (m mystring) toUpperCase() string {
	s := string(m)
	return strings.ToUpper(s)
}


/***************************************************************************************/

func main(){
	c := circle{2.0}
	r := rectangle{2.0,3.0}
	fmt.Println("area of circle c is ", c.Area())
	fmt.Println("area of rectangle r is ", r.Area())

	// Verifying whether modifying the value works
	(&c).changeRadius(4.0)  // c.changeRadius() also works. Thank you golang
	fmt.Println("area of circle c is ", c.Area())

	/* 
	Good points to know:
	1. Promoted fields work in methods
	2. If nested struct has a method implemented, then methods is also promoted to outer level (so cool)

	*/

	// ADVANCED STUFF : Methods on non structs
	// It can work on anything as long as it is a type

	str := mystring("abhishek")
	fmt.Println(str.toUpperCase())

	/* OUTPUT: 
	area of circle c is  12.566371
	area of rectangle r is  6
	area of circle c is  50.265484
	ABHISHEK
	*/
}