package main

import (
	"fmt"
	"math/rand"
	"time"
)

// stardater interface
// the time.Time type in the stdlib satisfies the stardater interface implicitly
type stardater interface {
	YearDay() int
	Hour() int
}

func stardate(t stardater) float64 {
	// YearDay returns what day of year it is
	// for example Jan 3rd being 3, Feb 2 being 33, Dec 31 being 365 or 366
	doy := float64(t.YearDay())
	// Hour() returns hour of the day (0 to 23)
	h := float64(t.Hour()) / 24.0
	// 1000 + day of year + (hour of the day / 24)
	return 1000 + doy + h
}

func main() {
	day := time.Date(2021, 12, 31, 18, 0, 0, 0, time.UTC)
	fmt.Printf("%.1f Curiosity has landed\n", stardate(day))

	mday := sol(1422)
	fmt.Printf("%.1f is a date in marsian notation\n", stardate(mday))
}

// due to the flexibility of interfaces
// we can now create a type for mars days
// and as long our new type satisfies the stardater interface
// we can pass it to the function stardate

// called sol because // https://en.wikipedia.org/wiki/Sol_(day_on_Mars)
type sol int

func (m sol) YearDay() int {
	return int((m % 668))
}

func (m sol) Hour() int {
	// returns a random number between 0 and 23
	return rand.Int() % 24
}
