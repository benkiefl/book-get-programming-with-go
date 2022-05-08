package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	for i := 10; i >= 0; i-- {
		fmt.Println(i)
		if rand.Intn(100) == 50 {
			fmt.Println("Launch failed")
			break
		}
	}

}
