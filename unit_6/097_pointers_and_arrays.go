// where would one use array(s) and pass it around with pointers
// instead of using slices to begin with?

// arrays have a fixed size, if dealing with a data structure with fixed size
// it can make sense to do it this way, for example with a chess board

package main

import "fmt"

type board [8][8]rune

func main() {

	var b board

	reset_board(&b)

	print_board(&b)

	wipe_board(&b)

	print_board(&b)

}

func reset_board(b *board) {
	wipe_board(b)
	b[0] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}
	b[1] = [8]rune{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'}
	b[6] = [8]rune{'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p'}
	b[7] = [8]rune{'r', 'n', 'b', 'q', 'k', 'b', 'n', 'r'}

}

func wipe_board(b *board) {
	for i := 0; i < len(b); i++ {
		for l := 0; l < len(b[i]); l++ {
			b[i][l] = '_'
		}
	}
}

func print_board(b *board) {
	for i := range b {
		for l := range b[i] {
			fmt.Printf("%c ", b[i][l])
		}
		fmt.Println()
	}
}
