package main

import (
	"fmt"
	"strings"
)

// Type switch example
func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("i stored string", strings.ToUpper(i.(string)))
	case int:
		fmt.Println("i strored int ", i)
	default:
		fmt.Println("i strored something else", i)
	}

}

func advancedInterface() {
	explain("hello")
	explain(10)
	explain(false)

	// i stored string HELLO
	// i strored int  10
	// i strored something else false

}
