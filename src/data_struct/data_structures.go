package main

import (
    "fmt"
    "reflect"
    )

func main() {
    //slices : like python list slicing

        letters := [] string{"a","b","c"}
        fmt.Println(reflect.TypeOf(letters))

        //or we can use make function to create a slice
         s := make([]string, 3)
         fmt.Println("emp:", s)
         s[0] = "a"
         s = append(s, "e", "f")
         fmt.Println("apd:", s)

    //Maps
        m:= make(map[string]int)

        m2 := map[string]int{"foo":1,"bar":2}

        fmt.Println("m2 is ", m2)

        m["k1"] = 123
        m["k2"] = 1234
        fmt.Println("m is of the type" ,reflect.TypeOf(m))

        fmt.Println(m)

        delete(m,"k2")
        fmt.Println(m)

        _,  k2_present := m["k2"]
        _, k1_present := m["k1"]
        fmt.Println("k1 and k2 are present or not ? ", k1_present, k2_present)

}