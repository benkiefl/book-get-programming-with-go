package main

import (
	"fmt"
	"math"
)

// get_Pi simply returns math.Pi
func get_Pi() float64 {
	return math.Pi
}

func main() {

	// create variable of type func() float64
	var Pi_1 func() float64

	// assign get_Pi to Pi_1
	Pi_1 = get_Pi

	// print the function address in memory
	fmt.Println(Pi_1)
	fmt.Println(get_Pi)
	// print the return value of Pi_1
	fmt.Println(Pi_1())

	// you can think of function as being the function itself
	// while function(), the () executes the function
	// probably not true in use of correct nomenclature, but for understanding's sake
}
