package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var year int = rand.Intn(2400)
	var month int = rand.Intn(12) + 1
	var day int

	switch month {
	case 2:
		if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
			day = rand.Intn(29) + 1
		} else {
			day = rand.Intn(28) + 1
		}
	case 4, 6, 9, 11:
		day = rand.Intn(31) + 1
	default:
		day = rand.Intn(30) + 1
	}

	fmt.Printf("year / month / day: %4d / %.2d / %.2d\n", year, month, day)
}
