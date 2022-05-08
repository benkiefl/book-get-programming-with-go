// methods frequently receive pointers to structures
// the panic is caused when the pointer is used in the method
// (simply calling a method with a nil pointer doesn't cause the panic)
// so you can guard against this from within the method of course by checking if pointer != nil

package main

import "fmt"

type person struct {
	age int
}

func (p *person) birthday() {
	if p == nil {
		return
	}
	p.age++
}

func main() {

	var p *person

	fmt.Println(p)

	// p.birthday() will simply return if p is a nil pointer
	p.birthday()
}

// how to handle nil in Go is up to you
// protect against it, blanket return, return 0, print an error, et cetera
