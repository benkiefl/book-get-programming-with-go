package main

import "fmt"

// you might have various rovers, cars, locations, etc
// so you can simply create a type
type rover struct {
	lat  float64
	long float64
}

func main() {

	var spirit rover
	spirit.lat = -32.43
	spirit.long = -32.332

	// composite literal for initializing structure
	opportunity := rover{lat: 32.54, long: 34.521}
	// the order of the fields doesn't matter
	// will work if struct has more fields as well

	// composite literal for initializing structure (2)
	curiosity := rover{34.543, 278.32}
	// here the number of fields has to be machted
	// and the order is important
	// only use with stable structs that don't change

	fmt.Println(spirit, opportunity, curiosity)
	// you can also print the field names along the way by using +v
	fmt.Printf("%+v\n", spirit) // {lat:-32.43 long:-32.332}

	// struct values are copied
	sunrise := curiosity
	sunrise.lat += 0.5
	fmt.Println(curiosity, sunrise)
	// this means struct (just as arrays) will be copied if passed to function
}
