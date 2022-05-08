// rot13 with spanish language string

package main

import "fmt"

func main() {

	s := "Hola, Estación Espacial Internacional"

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
