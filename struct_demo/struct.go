package main

import (
	"fmt"
)

/* 
Structs are equivalent of classes in other languages.
We can simply create a struct using `type alias`
Eg:
type StructName struct {
	field1 filedType1
	filed2 fieldType2
}
*/

//Employee struct holds employee information
type Employee struct {
	FirstName , LastName string // These are exported fields from this struct
	salary int
	fullTime bool
}

/**********************************************************/

//Nested structs
type Salary struct {
	basic int
	allowance int
	bonus int
}

type NestedEmployee struct {
	name string
	age int
	Salary // This is called promoted field. we can declare it like salart Salary too
}

/**********************************************************/
//Example of using functions as a field and using interface as a type
type fullnameType func(string,string) string

type Salaried interface{
	getSalary() int
}

type nestedInterfaceEmp struct{
	firstname, lastname string
	FullName fullnameType
	salary Salaried
}

func (s Salary) getSalary() int{
	return s.basic + s.allowance + s.bonus
}
/**********************************************************/


func sampleStruct() {
	//Initializing a struct directly. Instead we could have done something like below:
	/*
	var abhishek Employee; 
	abhishek.FirstName = "abhishek"
	....
	....
	*/

	abhishek := Employee{
		FirstName : "Abhishek",
		LastName : "Sharma",
		salary: 10,
		fullTime: true,
	}

	/*
	Or, we can provide it in sequential order too
	abhishek := Employee("abhishek", "sharma", 10 ,true)
	*/

	fmt.Println(abhishek.FirstName)

	// Anonymous struct: Without explicitly using type struct. Not used much as we have to 
	// use type again and again
	tommy := struct {
		name string
		breed string
	}{"Tommy", "Golden Retriever"}

	fmt.Printf("Type of tommy is : %T \n", tommy)
	fmt.Println(tommy.name, tommy.breed)
	
	/* Output is: 
	Abhishek
	Type of tommy is : struct { name string; breed string } 
	Tommy Golden Retriever
	*/

	joban := NestedEmployee{
		name: "Joban",
		age: 25,
		Salary: Salary{1000,200,300},
	}

	fmt.Println(joban.name, joban.basic) // Salary field are accessible directly to joban now, instead of doing joban.salary.basic
	// Take care in case of same fields inside and outside. Inside one are not promoted to upper level in case of a conflict

}

func nestedInterfaces(){
	aditya := nestedInterfaceEmp{
		firstname: "Aditya",
		lastname: "Bansal",
		FullName: func(firstname string,lastname string) string {
			return firstname + " "+ lastname
		},
		salary: Salary{100,200,300},
	}

	fmt.Println("aditya's salary is ", aditya.salary.getSalary())
	
	//Functions as field
	fmt.Println("fullname is ", aditya.FullName(aditya.firstname,aditya.lastname))

}



func main(){
	sampleStruct()
	nestedInterfaces()

}