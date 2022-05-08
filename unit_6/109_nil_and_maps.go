package main

import "fmt"

func main() {

	// same as with slices
	// declared without composite literal or built-in make and map has value of nil

	var nil_map map[string]string

	empty_map := map[string]string{}

	empty_map_1 := make(map[string]string)

	fmt.Println(nil_map == nil, empty_map == nil, empty_map_1 == nil)
	fmt.Println(nil_map, empty_map, empty_map_1)

	// maps can be read even if nil
	key, check := nil_map["keysearch"]
	if check {
		fmt.Println(key)
	}
	fmt.Println(key, check)
	fmt.Printf("%T %T\n", key, check)

	// writing to a nil map however panics
	// panic: assignment to entry in nil map
	// nil_map["hello"] = "world"

	// so if a function only reads from a map, you can pass a nil map to the function
	read_from_map(nil_map)
}

func read_from_map(m map[string]string) {
	fmt.Println(m["s"])
}
