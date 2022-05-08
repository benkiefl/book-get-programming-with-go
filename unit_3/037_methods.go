// methods are like functions with the difference of
// being associated to a type by way of a receiver before the method name
// so what does that mean???

// let's write a function and a method which do the same thing

package main

import "fmt"

type kg float64
type lbs float64

func convert_lbs_to_kg(l lbs) kg {
	return kg(l * 0.4535924)
}

// the following is a method, think of it as a kind of function with caveats
// type lbs float64 ; redundant, but important context
func (l lbs) lbstokg() kg {
	return kg(l * 0.4535924)
}

func main() {

	var l lbs = 205.32
	var k kg = convert_lbs_to_kg(l)
	var k1 = l.lbstokg()

	fmt.Println(l, k, k1)
}
