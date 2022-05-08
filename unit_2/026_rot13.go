package main

import "fmt"

func main() {

	const SHIFTVAL = 13

	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {

		c := message[i]

		if c != ' ' {
			c = message[i] + 13
		}

		if c > 'z' {
			c -= 26
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
