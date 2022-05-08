package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var total float64
	var value_coin float64

	for {

		switch coin_type := rand.Intn(6); coin_type {
		case 0:
			value_coin = 0.01
			total += value_coin
		case 1:
			value_coin = 0.02
			total += value_coin
		case 2:
			value_coin = 0.05
			total += value_coin
		case 3:
			value_coin = 0.10
			total += value_coin
		case 4:
			value_coin = 0.20
			total += value_coin
		case 5:
			value_coin = 0.50
			total += value_coin
		}

		fmt.Printf("added € %.2f coin, new total %2.2f\n", value_coin, total)

		if total > 20.00 {
			break
		}

		// alternatively you can count cents, use int type to add everything up
		// and for result in € do total / 100
		// possibly with some form of type conversion so go doesn't complain
		// above version seems accurate though

	}

}
