package main

import "fmt"

func main() {

	// I think it's much simpler to write

	var board [4][3]int

	for i := range board {
		for x := range board[i] {
			// this makes it so much more obvious what's happening
			board[i][x] = i
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

	// hint: if you have trouble! pen and paper
	// write the for loop down, draw the board, try to really get it

}
