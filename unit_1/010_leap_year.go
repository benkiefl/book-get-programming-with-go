package main

import "fmt"

func main() {

	var year int = 1997

	if ((year%4 == 0) && (year%100 != 0)) || (year%400 == 0) {
		fmt.Println("year", year, "is a leap year")
	} else {
		fmt.Println("year", year, "isn't a leap year")
	}

}
