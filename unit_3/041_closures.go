// this is taken from https://golangdocs.com/closures-in-golang
// much simpler example than in the book

package main

import "fmt"

func main() {

	i := 0

	// function literal (an anonymous function)
	// has no name ; is a closure
	// closure means it keeps references to outside values
	// without type inference it would be: var f = func() {}
	f := func() {
		// function has access to i
		i++
	}
	fmt.Println(i) // prints 0
	f()            // will increment i
	fmt.Println(i) // prints 1

	i = 0
	// you can even declare a function and immediately use it
	// without assigning it to a variable
	fmt.Println(i)
	func() {
		i++
	}()
	fmt.Println(i)

}
