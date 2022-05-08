package main

import "fmt"

type report struct {
	sol
	temperature
	humidity
	location
}

type sol int

type temperature struct {
	high, low celsius
}

type humidity struct {
	high, low float64
}

type location struct {
	lat, long float64
}

type celsius float64

// let's say we have a name collision
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func (h humidity) average() float64 {
	return (h.high + h.low) / 2
}

// let's create an average() method on the report type
func (r report) average() float64 {
	return (r.humidity.average())
}

func main() {

	report := report{
		sol:         5,
		temperature: temperature{high: 23, low: -10},
		humidity:    humidity{high: 87, low: 65},
		location:    location{lat: 59.94885521242605, long: 30.31703316085591},
	}

	// normally go would report an ambiguous selector report.average
	// our average() method on report type takes precedence over the others
	fmt.Printf("%+v\n", report.average())
	// no more ambiguous selector

	// languages like C++, Java, PHP, Python, Ruby, Swift can use composition
	// but they also have something called inheritance
	// Go's creators thought composition was more flexible, easier to reuse and change
}
