package main

import "fmt"

func main() {

	var s string      // default zero string ""
	var c rune = 'a'  // encoded as int32
	var c1 = 'b'      // from single quotes Go will infer a rune
	var c2 byte = 'c' // byte is an alias for uint8
	// old ASCII encoding (now subsect of Unicode): 128-character set
	// so byte (uint8) is the perfect variable type to encode ASCII characters

	fmt.Printf("%T %T %T\n", c, c1, c2)
	fmt.Println(c)
	fmt.Printf("%c\n", c)

	fmt.Println(s)

	fmt.Println(`raw string can be multiline /n/t/a tab empty lines etc all work`)

	// strings themselves can't be changed (are immutable, same in python, java, javascript)
	// strings can be accessed via index

	s = "starship enterprise"
	// s[] acces single byte ascii character
	fmt.Printf("%c%c%c%c\n", s[3], s[2], s[1], s[0])

	for i := 0; i < 19; i++ {
		fmt.Printf("%c ", s[i])
	}
	fmt.Println()
}
