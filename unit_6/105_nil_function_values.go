package main

import "fmt"

func main() {

	// a variable declared as a function type has the default value <nil>
	var fn func()
	fmt.Printf("%T %t\n", fn, fn == nil)

	// fn is a variable of type func() but no function is assigned, therefore defaults to <nil>
	// calling fn() results in a nil pointer dereference and leads to a panic
	// fn()

	// if fn were assigned a function, it wouldn't

	fn = function

	// now fn is assigned a function, and points to a memory address (therefore isn't nil)

	fmt.Printf("%T %t\n", fn, fn == nil)
	fmt.Println(fn)

	// no panic when calling fn()

	fn()
}

func function() {}
