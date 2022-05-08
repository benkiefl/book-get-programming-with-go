/* I came across a really nasty bug in 152_life_on_mars_v2.go which almost cost
me my sanity. It was a shadowed variable.
*/

package main

import "fmt"

func main() {

	n := 0
	fmt.Println(n)

	for i := 0; i < 1; i++ {
		n := 5
		fmt.Println(n)
	}

	fmt.Println(n)

}

/* this is legal Go code since n := 5 in the for scope is a new variable which
is specific to that scope. So this does not give a warning. This can be a nasty
source for bugs. See also
https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
*/
