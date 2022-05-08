package main

import (
	"fmt"
	"strconv"
)

func main() {

	// rune, byte, to string can easily be converted
	// remember rune and byte are simply aliases for int32 and uint8

	var a_rune rune = 'T'
	var b_rune rune = 34
	var a_byte byte = '2'
	var b_byte byte = 124

	fmt.Println(a_rune, b_rune, a_byte, b_byte)
	fmt.Println(string(a_rune), string(b_rune), string(a_byte), string(b_byte))

	// strconv.Itoa
	var i1 int = 32
	var s1 string = "string conversion Integer " + strconv.Itoa(i1) + " to ASCII"

	fmt.Println(s1)

	// Sprintf ; returns a string rather than displaying it
	s1 = fmt.Sprintf("let's try Sprintf %d here", i1)
	fmt.Println(s1)

	// strconv.Atoi
	// has a built-in error message, returns value + error code
	d1, err := strconv.Atoi("32")
	fmt.Println(d1, err)

	// a simple check could be
	d1, err = strconv.Atoi("4agfz")
	fmt.Println(d1) // 0
	if err != nil {
		fmt.Println(err) // strconv.Atoi: parsing "4agfz": invalid syntax
	}
}
