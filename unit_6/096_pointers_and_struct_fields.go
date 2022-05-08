package main

import (
	"fmt"
)

type character struct {
	name  string
	kind  string
	stats stats
}

type stats struct {
	level     int
	endurance int
	health    int
}

// levelUp takes a pointer to a stats struct as a parameter
func levelUp(s *stats) {
	s.level++
	s.endurance = 20 + (14 * s.level)
	s.health = 50 + s.endurance
}

func main() {

	char0 := character{
		name: "John Doe",
		kind: "Wizard",
	}

	// we can pass the memory address from the "stats struct within our char0 struct"
	levelUp(&char0.stats)

	fmt.Println(char0)

}
