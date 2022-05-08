package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var s string
	var s1 string
	var s2 string

	s = "国安部"
	s1 = "ФСБ"
	s2 = "CIA"

	// unicode/utf8 package
	// method DecodeLastRuneInString has two return values
	// first char of string + unicode value

	char, size := utf8.DecodeLastRuneInString(s)
	char1, size1 := utf8.DecodeLastRuneInString(s1)
	char2, size2 := utf8.DecodeLastRuneInString(s2)

	for _, c := range s {
		fmt.Printf("%c\t%d\t0x%.4x\n", c, c, c)
	}

	// _ (underscore) can be used a blank identifier if the index isn't needed
	// range can be used to iterate over a variety of collections

	for _, c := range s1 {
		fmt.Printf("%c\t%d\t0x%.4x\n", c, c, c)
	}

	for _, c := range s2 {
		fmt.Printf("%c\t%d\t0x%.4x\n", c, c, c)
	}

	// remember before to showcase ASCII we explicity used a byte / uint8
	// by default encoding will be 32bit (utf-8)
	fmt.Printf("%c\tbtyes: %d\t%T\n", char, size, char)
	fmt.Printf("%c\tbytes: %d\t%T\n", char1, size1, char1)
	fmt.Printf("%c\tbytes: %d\t%T\n", char2, size2, char2)

}
