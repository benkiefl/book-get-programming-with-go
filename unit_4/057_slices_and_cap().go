package main

import "fmt"

func main() {

	ar := [...]string{"a", "b", "c", "d", "e", "f", "g", "h"}
	sl := ar[2:5]

	fmt.Printf("%v %v\n", len(sl), cap(sl)) // 3 6

	// https://golang.org/pkg/builtin/#cap
	// cap(slice): the maximum length the slice can reach when resliced;
	// cap(sl) is 6 here because you can append elements up to max size of array
	// so with sl := ar[2:5] ;;; ar[0], ar[1] are inaccessible

	fmt.Println(sl)

	sl = append(sl, "1", "2", "3")

	fmt.Println(sl)
	fmt.Println(ar)

	// sl := ar[2:5] is shorthand for sl := ar[2:5:cap(ar)]
}
