package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	// short declaration can be used in for, if, switch function declarations
	// variables are then bound by those scopes and can't be accessed outside of it
	// keeps the program clean

	for i := 0; i < 2; i++ {
		fmt.Println("only place i exists", i)
	}

	x := rand.Intn(2)

	if i := rand.Intn(2); x == i {
		fmt.Println("'x' and 'i' are the same:", x, i)
	} else {
		fmt.Println("'x' and 'i' aren't the same:", x, i)
	}

	switch command := 'y'; command {
	case 'x':
		fmt.Println("case is x")
	case 'y':
		fmt.Println("case is y")
	default:
		fmt.Println("default")
	}

}
