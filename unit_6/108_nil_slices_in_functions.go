// add_items takes a slice of type []string
// since it utilizes append (which works with nil)
// I can pass nil to add_items and the function still works

package main

import "fmt"

func main() {

	slice := add_items(nil)

	fmt.Println(slice)

}

func add_items(s []string) []string {
	return append(s, "a", "b", "c")
}

// whenever a function accepts a slice
// assure that a nil slice has the same behavior as an empty slice
