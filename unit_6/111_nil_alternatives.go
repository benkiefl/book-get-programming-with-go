// instead of checking for nil you could also do something like this

package main

import "fmt"

// our number struct has a variable valid of type bool
// we flip that value to true if our "number is properly initialized"
type number struct {
	value int
	// valid bool is unambigous (as compared to working with nil)
	valid bool
}

func newNumber(v int) number {
	return number{value: v, valid: true}
}

// now we can simply check if valid is set to true
func (n number) String() string {
	// if valid isn't true, the number isn't set
	if !n.valid {
		return "not set"
	}
	return fmt.Sprintf("%d", n.value)
}

func main() {

	n := newNumber(42)
	fmt.Println(n)

	e := number{}
	fmt.Println(e)

}
