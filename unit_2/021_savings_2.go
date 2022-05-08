package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var total uint16
	var coin_value uint16
	rand.Seed(time.Now().UnixNano())

	for {

		switch coin_type := rand.Intn(6); coin_type {
		case 0:
			coin_value = 1
			total += coin_value
		case 1:
			coin_value = 2
			total += coin_value
		case 2:
			coin_value = 5
			total += coin_value
		case 3:
			coin_value = 10
			total += coin_value
		case 4:
			coin_value = 20
			total += coin_value
		case 5:
			coin_value = 50
			total += coin_value
		}

		fmt.Printf("adding ¢ %02d ; new total € %2d,%02d\n", coin_value, total/100, total%100)

		if total >= 2000 {
			break
		}
	}
}
