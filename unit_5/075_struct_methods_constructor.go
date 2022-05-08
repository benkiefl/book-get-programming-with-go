package main

import "fmt"

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

// below is a constructor function
// __init__ (python), initialize (ruby), _construct() PHP
// in Go constructors are not a special language feature
// they are normal functions
// newLocation (NewLocation if you need to export it)

// function newLocation takes two coordinate structs (dms notation) and returns location struct (degrees)
// due to my lack of experience in OOP I fail to see how this is super special, maybe it just isn't
// or I will be able to along the way
func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'W', 'w', 'S', 's':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {

	// 59.93430593323544, 30.324560270088195
	// Кasaner Kathedrale St Petersburg
	// google maps right click on map gives you lat and long in degrees
	lat := coordinate{59, 56, 3.501, 'N'}
	long := coordinate{30, 19, 28.417, 'E'}
	fmt.Println("Казанский кафедральный собор:", lat.decimal(), long.decimal())

	// bradbury landing curiosity
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'},
		coordinate{137, 26, 30.12, 'E'})
	fmt.Println("bradbury mars rover landing: ", curiosity)
}
