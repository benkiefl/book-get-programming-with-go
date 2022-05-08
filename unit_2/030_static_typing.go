package main

import (
	"fmt"
	"math"
)

func main() {

	// go has static typing

	var a_int int = 5
	var a_float float64 = 5.5

	// even with type inference, types are assigned

	b_int := 5
	b_float := 5.5

	fmt.Printf("%T %T %T %T\n", a_int, a_float, b_int, b_float)

	// c := b_int + b_float
	// --- invalid operation: mismatched types int and float64

	c := b_int + int(b_float)
	fmt.Println(c)
	// works, but will truncate b_float; int(b_float) = 5

	// some languages (eg python ruby javascript) use dynamic typing and aren't as strict

	// in Go we need to use type conversion
	// you have to be careful, especially when converting 'down'; float64 to float32 etc

	var c_int int64 = 9340238249023423
	fmt.Println(uint8(c_int))
	// value assigened to c_int is too large for uint8
	// due to overflow this will print something most likely not intended (191 in this case)

	// to get info about max and min values for types, math package provides min/max constants
	fmt.Println(math.MaxUint8) // maximum value for uint8 is 255
}
