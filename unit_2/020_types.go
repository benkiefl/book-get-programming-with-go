package main

import (
	"fmt"
)

// most common integer types simply are int and uint

// int		bit-size depending on environment
// int8		-128 to 127
// int16	-32,768 to 32,767
// int32	-2,147,483,684 to 2,147483647
// int64
// uint		bit-size depending on environment
// uint8	0 to 255
// uint16	0 to 65535
// uint32	0 to 4,294,967,295
// uint64

// to avoid overflow errors, on 64-bit systems int can be enough
// though on 32-bit systems those would have to be declared as int64

func main() {

	var x int64
	var y int // depends on your system

	// int64 and int are stil different types
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%[1]d%[2]d%[1]d%[2]d\n", 1, 2)

	// CSS (cascading style sheet)
	// encodes colors on screen each with range of 0 to 255
	// perfect use of uint8
	// imagine you're storing a lot of colors sequentially
	// use of uint8 could achieve considerable memory savings

	var red, green, blue uint8 = 0, 141, 213

	fmt.Printf("color: #%02x%02x%02x\n", red, green, blue)
	// for alphabets in uppercase (A-F) do %X instead of %x

	// video on hexadecimal https://www.youtube.com/watch?v=9xbJ3enqLnA
	// 15 and Hexadecimal - Numberphile

	fmt.Printf("\n")

	var u_int_8 uint8 = 255
	fmt.Printf("%08b %08b\n", u_int_8, u_int_8+1)
	fmt.Printf("%x %x\n", u_int_8, u_int_8+1)
	fmt.Printf("%d %d\n", u_int_8, u_int_8+1)

	fmt.Printf("\n")

	var int_8 int8 = 127
	fmt.Printf("%08b %08b\n", int_8, int_8+1)
	fmt.Printf("%x %x\n", int_8, int_8+1)
	fmt.Printf("%d %d\n", int_8, int_8+1)

	// math.MinInt32 math.MaxUint8 etc can give you max/min values

}
