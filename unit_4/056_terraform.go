package main

import "fmt"

type cel_bodies []string

func (p cel_bodies) terraform() {
	for i, _ := range p {
		p[i] = "new " + p[i]
	}
}

func classify_planets(planets []string) {
	terrestrial := planets[:4]
	gas_giants := planets[4:6]
	ice_giants := planets[6:]

	for i := range terrestrial {
		terrestrial[i] = "ter: " + terrestrial[i]
	}

	for i := range gas_giants {
		gas_giants[i] = "gas: " + gas_giants[i]
	}

	for i := range ice_giants {
		ice_giants[i] = "ice: " + ice_giants[i]
	}

}

func main() {

	planets := cel_bodies{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	classify_planets(planets[:])
	planets.terraform()

	fmt.Println(planets)
}
