package main

import "fmt"

// passing variables around is prone to errors and can be tedious

func main() {

	// a structure can bundle variables together
	// even variables of different types
	var rover struct {
		lat   float64
		long  float64
		is_on bool
		rank  int64
	}

	rover.lat = 2.32
	rover.long = 3.21
	rover.is_on = true
	rover.rank = 8

	fmt.Printf("Our rover is positioned at %.2f/%.2f \nit's on? %v, and ranks its condition at %d of 10\n",
		rover.lat, rover.long, rover.is_on, rover.rank)

}
