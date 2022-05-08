package main

import "fmt"

func main() {

	message := "L fdph, L vdz, L frqtxhuhg"

	for i := 0; i < len(message); i++ {

		c := message[i]
		c_new := message[i]

		if c >= 'a' && c <= 'z' {

			c_new = c - 3

			if c_new < 'a' {
				c_new += 26
			}
		}

		if c >= 'A' && c <= 'Z' {

			c_new = c - 3

			if c_new < 'A' {
				c_new += 26
			}
		}

		fmt.Printf("%c", c_new)

	}

}
