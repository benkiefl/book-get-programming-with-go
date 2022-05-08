package main

import "fmt"

/*
func function_name(parameters) return values {
	function body
}
*/

// lowercase function names: to be used in current package only
// Uppercase (first letter): to be used in other packages as well
// package.Function ; eg fmt.Printf

// kelvintocelsius converts kelvin to celsius
func kelvintocelsius(k float64) float64 {
	k -= 273.15
	// returns value to caller, where it's used to initialize a variable
	// the function has no outside effect (no side effect)
	// this is called by pass by value
	return k

}

// celsiustofahrenheit converts celsius to fahrenheit
func celsiustofahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

// kelvintofahrenheit converts kelvin to fahrenheit
// by reusing kelvintocelsius and celsiustofahrenheit
// for demonstration purposes
func kelvintofahrenheit(k float64) float64 {
	c := kelvintocelsius(k)
	f := celsiustofahrenheit(c)
	return f
}

//

func main() {

	var kelvin float64 = 294

	celsius := kelvintocelsius(kelvin)
	fahrenheit := celsiustofahrenheit(celsius)
	fahrenheit_2 := kelvintofahrenheit(kelvin)

	fmt.Printf("%.2f\t%.2f\t%.2f\t%.2f\n", kelvin, celsius, fahrenheit, fahrenheit_2)
}

// some functions are quite special, like fmt.Println
// it can accept one, two, or more parameters, can also accept various types
/*
func Println(a ...interface{}) (n int, err error)
*/
// this is called a variadic function
// type of parameter a is interface{}, known as the empty interface
// more about this later on
