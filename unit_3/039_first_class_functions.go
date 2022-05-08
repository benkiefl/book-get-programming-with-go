// this chapters suffers majorly from over-complicated examples
// and lack of detailed and clear explanation (imo)

// in go functions are first-class
// they work everywhere integers, strings, other types work

package main

import "fmt"

// do_x_times executes a function passed to it x times
// so functions can simply be passed as arguments to one another
func do_x_times(x int, function func()) {
	for i := 0; i < x; i++ {
		function()
	}
}

// prints a single line
func print_line() {
	fmt.Println("test")
}

func main() {

	// we're calling do_x_times with arguments 5 and function print_line
	do_x_times(5, print_line)

}
