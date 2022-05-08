package main

import (
	"fmt"
	"math"
)

type gps struct {
	current     location
	destination location
	world       world
}

type location struct {
	lat, long float64
	name      string
}

type world struct {
	radius float64
}

func (l location) description() string {
	return fmt.Sprintf("%v located at %.2f %2f", l.name, l.lat, l.long)
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p1.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (g gps) distance() float64 {
	return g.world.distance(g.current, g.destination)
}

func (g gps) message() string {
	return fmt.Sprintf("%.2f km remain to %s", g.distance(), g.destination.description())
}

type rover struct {
	gps
}

func main() {

	gps_mars := gps{
		current:     location{name: "Bradbury", lat: -4.5895, long: 137.4417},
		destination: location{name: "Elysium", lat: 4.5, long: 135.9},
		world:       world{radius: 3389.5},
	}

	fmt.Println(gps_mars)

	curiosity := rover{gps_mars}

	fmt.Println(curiosity.message())

}
