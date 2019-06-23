package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name   string `json:name`
	Salary int    `json:salary`
}

func MustMarshal(data interface{}) []byte {
	out, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	return out
}

func main() {

	var emp Employee
	data := []byte(`
		{
			"name":"Abhishek Sharma",
			"salary":10000
		}
		`)

	err := json.Unmarshal(data, &emp)
	if err == nil {
		fmt.Print(emp.Name)
	}

	mapD := map[string]int{"nameid": 1, "salary": 7}
	mapB := MustMarshal(mapD)
	fmt.Println(string(mapB))

}
