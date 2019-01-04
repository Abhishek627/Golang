package main

import "fmt"

func add (a int , b int) int {
    return a+b
}


func sums(nums ...int) {
//varaidic functions
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main(){
 //range: iterator for variety fo data structs
    nums := []int{1,2,3}

    sum:=0

    for temp ,num := range nums {
        fmt.Println(temp)
        sum+= num
    }

    fmt.Println(sum)

    //iterating over a map
    kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    //iterating over just keys
    for k := range kvs {
        fmt.Println("key:", k)
    }

    for i, c := range "go" {
        fmt.Println(i, c)
    }

    fmt.Println(add(5,5))
    sums(1, 2)
    sums(1, 2, 3)

    temp := []int{1,2,3,4,5}
    sums(temp...)


}

