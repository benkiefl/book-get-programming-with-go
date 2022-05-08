/*
	so the gist of it is that we can use structs with methods to do a form of OOP almost
*/

package main

import (
	"fmt"
	"math"
)

// world type to declare worlds
type world struct {
	radius float64
}

// location type, location in lat long (degrees)
type location struct {
	lat, long float64
}

// distance method for world type, takes two locations of type location
// and returns distance (float64) between the two locations
// utilizies our function rad to convert degrees to radians where needed
// calculation based on http://movable-type.co.uk/scripts/latlong.html
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {

	// we declare a new world using world type
	mars := world{radius: 3389.5}

	// we create locations using location type
	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}

	// we use the distance method on world type to get the distance and store result in dist
	dist := mars.distance(spirit, opportunity)

	// print dist variable (saved distance from before)
	fmt.Printf("%.2f km\n", dist)

}
