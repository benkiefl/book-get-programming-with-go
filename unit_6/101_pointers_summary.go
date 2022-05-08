package main

import "fmt"

func main() {

	var i int
	var p *int
	p = &i

	// the address operator & provides the memory address of a variable
	fmt.Println(&i)
	// pointers store memory addresses
	fmt.Println(p)
	// a pointer can be dereferenced with *
	fmt.Println(*p)

	// pointers allow mutation of values across function and method boundaries

	// pointers are most useful with structs and arrays
	// maps and slices use pointers behind the scenes
	// interior pointers can point at fields inside structures

	// use pointers when it makes sense, but don't overuse them
}
