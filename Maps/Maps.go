package main

import "fmt"

type info struct {
	Name    string
	Address string
}

func main() {
	// Maps structure is like map[key]value
	// In map all of the key must be of same data type and value must be of same data type
	// Map is indexed by set of unique key of same data type
	// In map we can have key of different data type and value of different data type
	// Simply saying what map does is that it maps keys to values

	// Empty map example
	emptyMap := make(map[string]int)

	fmt.Println("emptyMap: ", emptyMap)

	// Map is unordered
	mapWithSize := make(map[string]int)
	mapWithSize["test1"] = 1
	mapWithSize["test2"] = 2
	mapWithSize["test3"] = 3
	mapWithSize["test4"] = 4
	mapWithSize["test5"] = 5

	fmt.Println("\n mapWithSize: ", mapWithSize)

	//Maps can be created using builtin make function
	map1 := make(map[string]int)
	map1["dog"] = 42
	map1["cat"] = 40

	fmt.Println("\n map1: ", map1)

	// Map can contain structs as value
	//This type of map is called literal map
	map2 := map[string]info{
		"Roll01": {
			Name:    "Roshan Rana Bhat",
			Address: "Naxal",
		},
	}

	fmt.Println("\n map2: ", map2)

	// map also can be made using builtin make
	map3 := map[string]info{}
	map3["Roll02"] = info{
		Name:    "Tom Holland",
		Address: "USA",
	}

	fmt.Println("\n map3: ", map3)

	// Length of map
	fmt.Println("\n Length of mapWithSize is: ", len(mapWithSize))

	//delete map elements
	delete(mapWithSize, "test2")
	fmt.Println("\n mapWithSize after deleting element: ", mapWithSize)

}
