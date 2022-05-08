package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var count_1 float64 = 0
	var count_2 float64 = 0

	for i := 0; i < 100000000; i++ {
		var x int = rand.Intn(100)
		if x == 50 {
			count_2++
		}
		count_1++
	}

	var percentage float64 = count_2 / count_1

	fmt.Println(percentage) // should be ~1%
}
