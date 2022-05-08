package main

import "fmt"

func main() {

	const SHIFTVAL = 3

	// ABCDEFGHIJKLMNOPQRSTUVWXYZ
	// ABCdefGHIjklMNOpqrSTUvwxYZ
	// abcdefghijklmnopqrstuvwxyz

	alphabet := "ABCdefGHIjklMNOpqrSTUvwxYZ"

	for i := 0; i < 26; i++ {
		c := alphabet[i]
		c_new := c + SHIFTVAL

		if c <= 'z' && c_new > 'z' {
			c_new -= 26
		} else if c >= 'A' && c <= 'Z' && c_new > 'Z' {
			c_new -= 26
		}
		fmt.Printf("%c", c_new)
	}
	fmt.Println()
}
