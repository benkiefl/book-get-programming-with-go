package main

import "fmt"

// you can create your own types
// the newly created type is called celsius
// it behaves like a float64 but is not an alias for it
// type conversion will need to take place
// this practice can make code more readable and avoid mistakes
type celsius float64
type fahrenheit float64

// since celsius and fahrenheit are now types behaving as float64
// but are separate types nonetheless
// eg adding fahrenheit to celsius value will throw an error

func main() {

	var temp_c celsius // far more descriptive than var temp_c float64
	var temp_f fahrenheit

	// numeric values are 'untyped'
	// means they can be assigned to many types of variables
	var f float64
	var i int64
	var b byte
	f = 2
	i = 2
	b = 2
	fmt.Printf("%T %T %T\n", f, i, b)

	temp_c = 1
	temp_f = 2

	// fmt.Println(temp_1 + temp_2)
	// invalid operation: mismatched types celsius and fahrenheit

	fmt.Printf("%T %T\n", temp_c, temp_f)
	// as you can see, the types are main.celsius and main.fahrenheit
	// refering to types happens via dot notation
	// similar to how functions from other packages are called eg fmt.Printf
}

// when you create your own types, of course you can use them in functions as well
// be it as parameters or return values
// but the functions will have to written accordingly
// remember: type celsius kelvin etc are not aliases for anything but their own type

/*
func kelvintocelsius(k kelvin) celsius{
	return celsius(k - 273.15)
}
*/

// regarding methods, classes I found this to be helpful
// http://www.golangpatterns.info/object-oriented/classes
