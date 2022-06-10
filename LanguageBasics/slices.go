package LanguageBasics

import (
	"fmt"
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

	e = append(e, 1, 2, 3)

	printSlice(`e:`, e)

	fmt.Printf("%v\n", e)

	return a
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
