// this program is meant to portray the complete execution flow
// using multiple defers, a panic and returning from a function that panicked

package main

import "fmt"

func main() {

	f()

	fmt.Println("will be printed even though f() panics")

}

func f() {

	// defer pushes onto a stack
	// Last function pushed onto the stack, first to be executed
	// LIFO (last in first out)

	defer fmt.Println("A")
	defer fmt.Println("B")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("C", r)
		}
	}()

	defer fmt.Println("D")

	panic("PANIC")
}

/*	Summary
Errors are values
Great deal of flexibility when dealing with Errors
Custom error types are possible (by satisfying the error interface)
defer keyword helps to clean up before function returns
type assertions can convert an interface to a concrete type or another interface
do not panic - return an error instead
*/
