/*	Error handling in other languages
exceptions, can be thought of as opt in
involve specialized code, specialized keywords (try, catch, throw, raise, rescue, ...)
*/

/*	Error handling in Go
Go doesn't have exceptions, but pushes developers to consider errors
Error values don't require spcialized keywords
Errors are simple values -> are simpler, and more flexible
Panics in Go should be rare (division by zero leads to panic)
*/

package main

import "fmt"

func main() {

	// defer pushes the function onto a stack
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
	}()

	var zero int
	_ = 23 / zero // division by zero, causes panic
	// the panic causes a return

	fmt.Println("This statement is not reachable")
}
