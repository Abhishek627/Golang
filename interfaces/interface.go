package main

import (
	"fmt"
)

//Inteface is collection of method signatures that an Object can implement
/*
In golang, there is no need to implicilty declare `implements` in the method implementing the interface, if the signature matches, 
type is said to be implementing the said interface
*/
// Interface can't implement an interface but we can create a new interface using other two


type shape interface {
	Area() float64 // Name_of_func ReturnType
	Perimeter() float64
}

/*
Interface in golang have 2 parts: (wtf?)
a. Static part --  refers to the interface itself
b. Dynamic part -- variable of interface holds the value of type that is implementing the interface
*/

// Implementing the interface itself using a struct.
//Should implement all the methods of an interface
type rect struct {
	width float64
	height float64
}

func (r rect)  Area() float64 {
	return r.height * r.width
}

func (r rect) Perimeter() float64 {
	return 2*( r.width+ r.height)
}

//Empty interfaces: can hold any value as all types implement it
func explainEmpty(i interface{}) {
	fmt.Printf("values given to this func is of type %T and value %v \n",i ,i)
}

//Implementing multiple interfaces
type printStat interface{
	Desc() string
}

//Now rect implements 2 interfaces, shape and printStat
func(r rect) Desc() string{
	return fmt.Sprintf("This shape as height and width of %f and %f ",r.height , r.width)
}

type walks interface{
	canWalk() bool
}

func main(){
	var s shape
	s = rect{2.0,3.0}
	fmt.Printf("Type of s is %T \n",s)
	fmt.Printf("value of s is %v \n",s)
	fmt.Println("area of rectangle s",s.Area())
	/*
	Type of s is main.rect 
	value of s is {2 3} 
	area of rectangle s 6
	*/

	explainEmpty(s)

	// Using type assertion to get underlying dynamic value of an interface using syntaz `i.(Type)`
	c , ok := s.(rect)
	//fmt.Println("stats of type rect is: ", s.Desc()) This will throw an error saying  s.Desc undefined (type shape has no field or method Desc)
	fmt.Println("stats of type rect is: %s and perimeter of this is %f", c.Desc(), c.Perimeter(), ok)

	// val := s.(walks) this will create a panic since s doesn't implements our walks interface
	val, present := s.(walks) // This is totally cool as present is false
	fmt.Print(val,present)  //returns <nil> false

}