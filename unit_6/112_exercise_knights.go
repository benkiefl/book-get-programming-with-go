package main

import "fmt"

type character struct {
	name  string
	hands *item
}

type item string

func (c *character) pickup(i *item) {
	fmt.Printf("%v picks up item %v\n", c.name, *i)
	c.hands = i
}

func (c *character) giveto(c1 *character) {
	if c.hands == nil {
		fmt.Printf("%v's hands are empty, can't give something to %v\n", c.name, c1.name)
	}
	fmt.Printf("%v gives item %v to %v\n", c.name, *c.hands, c1.name)
	c1.hands = c.hands
	c.hands = nil
}

func (c character) String() string {
	if c.hands == nil {
		return fmt.Sprintf("%v with an empty left hand", c.name)
	}
	return fmt.Sprintf("%v with a %v in his left hand", c.name, *c.hands)
}

func main() {

	arthur := character{name: "Arthur", hands: nil}
	knight := character{name: "Knight", hands: nil}
	weapon := item("sword")

	fmt.Println(arthur)
	fmt.Println(knight)

	arthur.pickup(&weapon)
	arthur.giveto(&knight)

	fmt.Println(arthur)
	fmt.Println(knight)

}
