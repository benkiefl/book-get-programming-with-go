package main

import "fmt"

// not all data structures are copied when assigned or passed as arguments
// maps for example are always passed by reference, via pointers, see 19.2

// so we could tell a function that it takes a pointer to a map
func demolish(m *map[string]string, s string) {
	// but that is verbose and completely unnecessary
	delete(*m, s)
}

// here we have the same function without explicitly expecting a pointer to a map
// maps are passed by reference, so pointers, automatically
func demolish2(m map[string]string, s string) {
	delete(m, s)
}

func main() {

	m := map[string]string{
		"foo":   "bar",
		"alice": "bob",
	}

	fmt.Println(m)
	demolish(&m, "foo")
	demolish2(m, "alice")
	fmt.Println(m)
}
