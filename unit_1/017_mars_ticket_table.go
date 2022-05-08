/*
	build program that generates 10 random tickets

	space agencies:
		国家航天局 (CNSA), NASA, ESA,
		Роскосмос (Roscosmos State Corporation),
		JAXA 国立研究開発法人宇宙航空研究開発機構
	distance: 62 100 000
	speed between: 16 - 30 km/s
	price:
		faster ships more expensive
		range: 36 to 50 million
		double price for round trips

*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const CNSA string = "CNSA"
	const NASA string = "NASA"
	const ESA string = "ESA"
	const RSC string = "RSC"
	const JAXA string = "JAXA"

	const DIST float64 = 62100000

	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Spaceline\tDays\tTrip type\tPrice [m]\n")
	fmt.Printf("=================================================\n")

	for i := 0; i < 10; i++ {

		switch space_agency := rand.Intn(4); space_agency {
		case 0:
			fmt.Printf("%s\t\t", CNSA)
		case 1:
			fmt.Printf("%s\t\t", NASA)
		case 2:
			fmt.Printf("%s\t\t", ESA)
		case 3:
			fmt.Printf("%s\t\t", RSC)
		case 4:
			fmt.Printf("%s\t\t", JAXA)
		}

		speed := float64(rand.Intn(15) + 16)
		trip_days := (DIST / speed) / 60 / 60 / 24

		fmt.Printf("%d\t", int(trip_days))

		trip_code := rand.Intn(2)

		if trip_code == 0 {
			fmt.Printf("one-way\t\t")
		} else {
			fmt.Printf("round-trip\t")
		}

		var price int = int(speed) + 20

		if trip_code == 1 {
			price *= 2
		}

		fmt.Printf("%d\n", price)

	}

}
