// anonymnous functions whenever you need to create a function on the fly
// for example when returning a function from another function

package main

import "fmt"

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	// we declare an anonymous function
	return func() kelvin {
		// remember, anonymous functions are closures
		// they keep references to the variables in the surrounding scope
		// so we can still do s() + offset
		return s() + offset
	}
	// the variables captured "survive"
}

func main() {
	var offset kelvin = 10
	sensor := calibrate(realSensor, offset)
	fmt.Println(sensor())
}
