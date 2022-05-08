package main

import "fmt"

// to embed a type in a structure specify the type without the field name
// all methods of temperature and location are now directly accessible through the report type
type report struct {
	sol         // automatically same name as type
	temperature // - " -
	location    // - " -
}

// you can embed any type in a structure, not just structures
type sol int

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {

	report := report{
		sol:         15,
		temperature: temperature{high: 23, low: -10},
		location:    location{lat: 59.94885521242605, long: 30.31703316085591},
	}

	fmt.Printf("%+v\n", report.average())

	// temperature and location are accessible just the same
	// (as if we had specified their field names explicitly)

	fmt.Printf("%+v %+v\n", report.temperature, report.location)

	// not only that, embedding allows use to access the fields of temperature
	// and location directly from within report

	fmt.Printf("%+v\n", report.temperature.high)
	// we can now just as easily write
	fmt.Printf("%+v\n", report.high)

}
