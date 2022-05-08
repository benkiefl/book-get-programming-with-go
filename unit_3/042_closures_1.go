package main

import "fmt"

func main() {

	i := 0
	j := 0

	fmt.Println(i, j)

	func() {
		i++
	}()

	fmt.Println(i, j)

}
