package main

import "fmt"

type character struct {
	name, power string
	age         int
}

func main() {

	wizard := character{
		name:  "Merlin",
		power: "frost spell",
		age:   158,
	}

	// by convenience composite literals are sometimes prefixed with the address operator
	wizard2 := &character{
		name:  "Merlin2",
		power: "whatever",
		age:   20,
	}

	// in case of wizard, we have to pass &wizard as an argument
	// in case of wizard2, we can simply do wizard2, wizard2 already is a pointer
	change_power(&wizard)
	change_power(wizard2)

	// struct pointers don't have to be explicitly dereferenced
	fmt.Printf("%s is the same as %s\n", (*wizard2).power, wizard2.power)
	// but you can't dereference a non-pointer ofc, so the following doesn't work
	// fmt.Printf("%s is the same as %s\n", (*wizard).power, wizard.power)
	// ./093_pointers_and_structs.go:33:39: invalid indirect of wizard (type character)

	// also if I were to do something like
	x := wizard  // x is a copy of wizard
	y := wizard2 // y is a pointer, pointing to the same sturct as wizard2 does
	// of course you can also just do
	z := &wizard // now x is a pointer to wizard

	fmt.Printf("%T %T %T\n", x, y, z)

}

func change_power(p *character) {
	// struct pointers don't have to be explicitly derferenced
	// instead of doing (*p).power = "magic flames" we can do
	p.power = "magic flames"
	// in Go function and method parameters are passed by value
	// by using pointers and dereferencing them we can change underlying data easily
	// like pass around a struct, it's never copied, we always work on "the original"
}
