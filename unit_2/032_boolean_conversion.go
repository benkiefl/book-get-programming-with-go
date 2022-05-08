package main

import "fmt"

func main() {

	// boolean to text

	yesno_bool := true
	yesno_text := fmt.Sprintf("%v", yesno_bool)

	fmt.Println(yesno_text)

	// text to boolean
	// you can assign the result of a condition directly to a variable

	yesno_text2 := (yesno_bool == true)
	fmt.Println(yesno_text2)

}
