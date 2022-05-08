package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	lat, long float64
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func main() {

	mars := world{radius: 3389.5}
	earth := world{radius: 6371.0}

	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 28, 21.5, 'E'})
	opportunity := newLocation(coordinate{1, 56, 46.3, 'S'}, coordinate{354, 28, 24.2, 'E'})
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0, 'E'})

	// hard coding this is ugly but solving this algorithmically is non-trivial
	fmt.Printf("distance between spirit and opportunity: %.2f\n", mars.distance(spirit, opportunity))
	fmt.Printf("distance between spirit and curiosity: %.2f\n", mars.distance(spirit, curiosity))
	fmt.Printf("distance between spirit and insight: %.2f\n", mars.distance(spirit, insight))
	fmt.Printf("distance between opportunity and curiosity: %.2f\n", mars.distance(opportunity, curiosity))
	fmt.Printf("distance between opportunity and insight: %.2f\n", mars.distance(opportunity, insight))
	fmt.Printf("distance between curiosity and insight: %.2f\n", mars.distance(curiosity, insight))

	st_petersburg := location{59.948855212426054, 30.31703316085591}
	moscow := location{55.752398821217604, 37.617473953166076}

	fmt.Printf("distance between petersburg and moscow: %.2f\n", earth.distance(st_petersburg, moscow))
}
