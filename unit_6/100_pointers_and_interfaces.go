package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

type frog struct{}

func (f frog) talk() string {
	return fmt.Sprintf("quak quak")
}

type lion struct{}

func (l *lion) talk() string {
	return fmt.Sprintf("roar roar")
}

func main() {

	f := frog{}
	l := lion{}

	// both f and &f satisfy the talker interface, as demonstrated with
	fmt.Printf("%s %s\n", shout(f), shout(&f))

	// however, if our method uses a pointer reciever, we have to explicitly use it with a pointer
	// shout(l)
	// cannot use l (variable of type lion) as talker value in argument to shout: missing method talk (talk has pointer receiver)
	fmt.Printf("%s\n", shout(&l))
}

func shout(t talker) string {
	return fmt.Sprintf(strings.ToUpper(string(t.talk())))
}

// in summation:
// value methods can be invoked on pointers and values
// pointer methods can only be invoked on pointers!
// see https://golang.org/doc/effective_go#pointers_vs_values
