package main

import "fmt"

func main() {
	fmt.Println("2 + 2 =", 2+2)
	fmt.Println("2 - 2 =", 2-2)
	fmt.Println("2 * 2 =", 2*2)
	fmt.Println("2 / 2 =", 2/2)
	fmt.Println("5 % 3 =", 5%3)

	var x int = 5
	// shorthand declaration: x := 5

	fmt.Println("x =", x)
	x++
	fmt.Println("x =", x)
	x--
	fmt.Println("x =", x)
	x += 5 // x = x + 5
	x -= 2 // x = x - 2
	x /= 2 // x = x / 2
	x *= 4 // x = x * 4
	fmt.Println("x =", x)

}
