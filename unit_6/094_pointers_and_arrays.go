package main

import "fmt"

func main() {

	// just as with structs, composite literals with arrays can be prefixed with the address operator (&)

	array := &[4]int{0, 1, 2, 3}

	// arrays also provide automatic dereferencing
	// (*array)[0] is effectively the same as array[0]

	fmt.Printf("%v %v %v %v\n", (*array)[0], (*array)[1], array[2], array[3])

	// composite literals for slices and maps can be prefixed with address operator (&)
	// but automatic dereferencing works only with structs and arrays

}
