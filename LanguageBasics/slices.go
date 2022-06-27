package LanguageBasics

import (
	"fmt"
	"strings"
)

// Slices are Dynamic Arrays

func SliceExample() []int {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var slice []int = primes[1:4]

	fmt.Println(`
	-------------------------
	---- Slice Example ------
	-------------------------
	`)

	fmt.Println(slice)

	slice = append(slice, 15)

	fmt.Println(slice)

	return slice
}

func MakeSlice() []int {
	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	a := make([]int, 5)

	fmt.Println(`
	------ Make Slice -------
	`)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	e := make([]int, 0)

	// We can add more than one element at a time.
	e = append(e, 1, 2, 3)

	printSlice(`e:`, e)

	fmt.Printf("%v\n", e)

	return a
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func TicTacToeMultiDimensional() {
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println(`
	------ MultiDimensional Slice -------
	`)

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
