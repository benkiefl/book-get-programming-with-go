package main

import (
	"fmt"
	"strings"
)

// more typically, interfaces are declared as named types
// by convention their names often end in -er
// describing the action the interface is about

type talker interface {
	talk() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

// if declaring an interface as a named type
// we can use that interface just where we would use types in general

// shout takes an interface of type talker as an argument
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {

	// variable m of type martian
	var m martian

	// martian type has a method talk()
	fmt.Println(m.talk())
	// martian type satisfies the interface talker

	// function shout() takes an argument of talker
	// this can be either a variable of a type which satisfies the interface
	shout(m)
	// or we can pass the type directly
	shout(martian{})

	// this means that ANY type which satisfies the talker interface
	// will work with function shout(t talker)
}
