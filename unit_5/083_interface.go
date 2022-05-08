package main

import (
	"fmt"
	"strings"
)

// the variable t is of type interface
// interfaces primarily are about what a type can do not what value it holds
// can hold any type that satisfies the interface
var t interface {
	talk() string
}

// any type implementing the method talk() in above way
// (taking no arguments and returning a string)
// is said to "satisfy" the interface

type martian struct{}
type laser int

// both types will have a method called talk() and return a string

func (m martian) talk() string {
	return "nack nack"
}

func (l laser) talk() string {
	return strings.Repeat("pew ", int(l))
}

func main() {

	// variable t of type interface can be of type martian or type laser
	// both types satisfy the interface by implementing talk() string

	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())

	// interaces therefore provide polymorphism (t can take many shapes)
}
