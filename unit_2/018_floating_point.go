// https://www.youtube.com/watch?v=PZRI1IfStY0
// link: Floating Point Numbers - Computerphile

package main

import (
	"fmt"
	"math"
)

func main() {

	const PI64 = math.Pi
	const PI32 float32 = math.Pi

	fmt.Println(PI64)
	fmt.Println(PI32)

	// in some cirtumstances (eg thousand of vertices in 3D games)
	// it may make sense to sacrifice precision for memory savings

	var i int64 // default value, called the zero value
	var f float64

	fmt.Printf("%d\n", i)       // 0
	fmt.Printf("%2.2f\n", f)    // 0.0
	fmt.Printf("%4.2f\n", 3.14) // % width.precision

	// floating point number rounding error

	third := 1.0 / 3.0
	fmt.Println(third + third + third)
	value := 0.1
	value += 0.2
	fmt.Println(value)

	// if mission critical you can store decimals as int
	// perform multiplication before division, decreases rounding errors

	celsius := 21.0
	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Println(fahrenheit) // 69.8

	fahrenheit = (9.0 / 5.0 * celsius) + 32.0
	fmt.Println(fahrenheit) // 69.80000000000001

	// due to rounding errors floating point numbers can't be compared directly

	value = 0.1
	value += 0.2

	fmt.Println(value == 0.3) // false

	// if you absolutely must:
	// use the math.Abs function to determine absolute difference
	// be sure to pick a good threshhold
	// floating point errors can accumulate rather quickly

	fmt.Println(math.Abs(value-0.3) < 0.0001) // true

}
