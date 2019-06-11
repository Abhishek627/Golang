package datastruct

import "fmt"

// ArrayDemo :Starting array demo now
func ArrayDemo() {
	fmt.Println("\n ***************** STARTING ARRAY DEMO NOW  *****************")
	var arr1 = [5]int{1, 2, 3, 4, 5}

	arr2 := [...]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)

	//Iterating an array

	for idx, val := range arr1 {

		fmt.Printf("idx: %d, val: %d \n", idx, val)

	}
	fmt.Println(" \n ***************** ENDING ARRAYS DEMO NOW  *****************")

}
