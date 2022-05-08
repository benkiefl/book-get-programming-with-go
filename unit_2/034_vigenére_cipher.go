package main

import "fmt"

func main() {

	// s := "ABC, DEF, GHI, JKL, MNO, PQR, STU, VWX, YZ"
	// k := "GG"
	s := "CSOITEUIWUIZNSROCNKFD"
	k := "GOLANG"

	for i, x := 0, 0; i < len(s); i++ {

		var c byte
		var shiftval byte

		if x == len(k) {
			x = 0
		}

		shiftval = k[x] - byte('A')

		if s[i] >= 'A' && s[i] <= 'Z' {
			c = s[i] - shiftval
			if c > 'Z' || c < 'A' {
				c += 26
			}
		} else {
			c = s[i]
		}

		fmt.Printf("%c", c)

		x++
	}

	fmt.Println()

}
