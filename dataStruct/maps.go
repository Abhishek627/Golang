package datastruct

import "fmt"

//Map :How to create a map and use it. Map is like dict in python
func Map() {
	// How to create a map and use it. Map is like dict in python
	fmt.Println(" ***************** STARTING MAPS DEMO NOW  *****************")

	//M1
	var map1 = make(map[string]int)
	map1["Hello"] = 1
	map1["Hi"] = 2

	//M2
	m2 := map[string]int{"foo": 1, "bar": 2}

	fmt.Println(map1) // map[Hello:1 Hi:2]
	fmt.Println(m2)

	for k, v := range map1 {
		println(k, v)
	}

	delete(map1, "Hello")

	id, p := map1["Hello"]
	print(id, p) // 0, false because we deleted hello from the map

	fmt.Println("\n ***************** ENDING MAPS DEMO NOW  *****************")

}
