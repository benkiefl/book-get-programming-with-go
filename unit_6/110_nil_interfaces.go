// %v with the # flag gives a go-syntax represenation of the value
// (by invoking the GoStringer interface), that way we can have a better look at our interface{}

package main

import "fmt"

func main() {

	// an interface variable declared without assignment is nil

	var i interface{}

	// interface type and value are both nil
	fmt.Printf("%T %v %#v %v\n", i, i, i, i == nil)

	// when a variable with an interface type is assigned a value
	var p *string
	i = p
	// the interface internally points to the type and value of that variable
	// both need to be nil for the variable to equal nil

	fmt.Printf("%T %v %#v %v\n", i, i, i, i == nil)
	// special attention to the output of %#v -> (*string)(nil)

	// if you are confused: the empty interface is satisfied by ANY type
	// which makes sense if you think about it

	i = nil

	fmt.Printf("%T %v %#v %v\n", i, i, i, i == nil)

	// for a beginner there is little real use for empty interfaces{} as variables
	// mostly used for marshaling/unmarshaling data or reflection
	// both reflection and empty interfaces{} should be avoided as a beginner
}
