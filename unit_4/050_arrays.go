package main

import "fmt"

func main() {

	ar0 := [...]int{
		0,
		1,
		2,
		3,
	}

	// assigning an array copies the values
	// same happens when passing it to a function
	// remember: pass by value
	ar1 := ar0

	// changes only ar1
	ar1[1] = 0

	// ar0[1] has 1 as value, ar1[1] has 0 as value
	fmt.Println(ar0[1], ar1[1])

	// [3]string and [4]string are both collection of strings
	// but they are different types

	string_array := [3]rune{'a', 'b', 'c'}
	var string_array_2 [4]rune

	// following won't work
	// string_array_2 = string_array

	var string_array_ [3]rune
	string_array_ = string_array

	fmt.Println(string_array, string_array_2, string_array_)

}
