// this is most likely an example of what not to do?!
// doesn't even use interfaces (which the assignment asks for)
// it's the first approach I came up with and felt the most "natural" to me
// I guess I don't yet understand interfaces well (as I don't see any need for them here)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type animal struct {
	name string
	food string
}

func (a animal) eat() {
	fmt.Printf("%v is eating %s\n", a.name, food(a))
}

func (a animal) move() {
	fmt.Printf("%v is %s\n", a.name, random_movement())
}

func main() {

	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		{name: "lion", food: "meat_eater"},
		{name: "gazelle", food: "plant_eater"},
		{name: "dog", food: "meat_eater"},
		{name: "elephant", food: "plant_eater"},
	}

	var random_animal int
	var random_action int
	hour := int(0)
	global_timer := 0

	for {
		time.Sleep(500 * time.Millisecond)
		if global_timer > 72 {
			break
		}

		fmt.Printf("%2d:00 ", hour)

		if day_check(hour) == true {
			random_animal = rand.Intn(4)
			random_action = rand.Intn(2)
		} else {
			fmt.Println("All the animals are sleeping")
			if hour >= 24 {
				hour = 0
			} else {
				hour += 1
			}
			global_timer += 1
			continue
		}

		switch random_action {
		case 0:
			animals[random_animal].eat()
		case 1:
			animals[random_animal].move()
		}

		hour += 1
		global_timer += 1
	}

}

func day_check(hour int) bool {
	if hour > 6 && hour < 21 {
		return true
	} else {
		return false
	}
}

func food(a animal) string {
	if a.food == "meat_eater" {
		return random_meat()
	} else if a.food == "plant_eater" {
		return random_plant()
	} else {
		return "error"
	}
}

func random_movement() string {
	action := rand.Intn(3)
	switch action {
	case 0:
		return "roaring"
	case 1:
		return "pacing"
	case 2:
		return "stomping"
	default:
		return "error"
	}
}

func random_meat() string {
	v := rand.Intn(3)
	switch v {
	case 0:
		return "veal"
	case 1:
		return "dog food"
	case 2:
		return "human flesh"
	default:
		return "error"
	}
}

func random_plant() string {
	v := rand.Intn(3)
	switch v {
	case 0:
		return "broccoli"
	case 1:
		return "banana"
	case 2:
		return "apples"
	default:
		return "error"
	}
}
