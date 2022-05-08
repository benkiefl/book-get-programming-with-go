package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type turtle struct {
	x, y int
}

func (t *turtle) up() {
	t.y -= 1
}

func (t *turtle) down() {
	t.y += 1
}

func (t *turtle) right() {
	t.x += 1
}

func (t *turtle) left() {
	t.x -= 1
}

// constants for our matrix size, row and column
const row = 5
const col = 10

// declare matrix using those constants
var matrix [row][col]rune

func draw(t *turtle) {

	// code so we wrap around if leaving our matrix
	if t.y < 0 {
		t.y += row
		t.x -= 1
	}
	if t.y >= len(matrix) {
		t.y -= row
		t.x += 1
	}
	if t.x < 0 {
		t.x += col
		t.y -= 1
	}
	if t.x >= len(matrix[0]) {
		t.x -= col
		t.y += 1
	}

	// here we initialize the matrix and at the same time print it
	for i := range matrix {
		for l := range matrix[i] {
			// for our given position store an 'x' in the matrix
			// all other positions store an '_' in the matrix
			if (i == t.y) && (l == t.x) {
				matrix[i][l] = rune('x')
			} else {
				matrix[i][l] = rune('_')
			}
			fmt.Printf("%c", matrix[i][l])
		}
		fmt.Println()
	}
}

func main() {
	// declare turtle with starting position
	t := turtle{0, 0} // 0, 0 corresponds to left top

	clear_screen()

	left_to_right(&t)
	bottom_to_top(&t)
	top_to_bottom(&t)
	right_to_left(&t)

	time.Sleep(3 * time.Second)

}

func clear_screen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func left_to_right(t *turtle) {
	for {
		clear_screen()
		draw(t)
		fmt.Println(t.x, t.y)
		if t.x == col-1 && t.y == row-1 {
			break
		}
		t.right()
		time.Sleep(100 * time.Millisecond)
	}
}

func bottom_to_top(t *turtle) {
	for {
		clear_screen()
		draw(t)
		fmt.Println(t.x, t.y)
		if t.x == 0 && t.y == 0 {
			break
		}
		t.up()
		time.Sleep(100 * time.Millisecond)
	}
}

func top_to_bottom(t *turtle) {
	for {
		clear_screen()
		draw(t)
		fmt.Println(t.x, t.y)
		if t.x == col-1 && t.y == row-1 {
			break
		}
		t.down()
		time.Sleep(100 * time.Millisecond)
	}
}

func right_to_left(t *turtle) {
	for {
		clear_screen()
		draw(t)
		fmt.Println(t.x, t.y)
		if t.x == 0 && t.y == 0 {
			break
		}
		t.left()
		time.Sleep(100 * time.Millisecond)
	}
}
