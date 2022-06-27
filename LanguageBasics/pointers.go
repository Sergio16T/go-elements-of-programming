package LanguageBasics

import "fmt"

/*
Why Pointers??

GoLang is always passByValue in order to provide the ability to update the reference to the value provided in a function,
we need to provide a memory address to the value so that we can update the value at that address

In Go this is accomplished with pointers

Instead of copying a large amount of data every time you need to pass it, programmers could pass its address.

Pointers provide shared access to a value.

NOTES:

1. If we want a method to modify the parameter provided outside the scope of the function, we must use a pointer!
2. If we have a large amount of data, we can pass a pointer to its location for better performance

& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.
* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
*/

func noPointer() string {
	return "string"
}
func pointerTest() *string {
	return nil // You cannot return nil from a function that returns a string
	// but you can return nil from a function that returns a pointer to a string!
	// The zero value of a pointer is nil
}
func pointerTestTwo() *string {
	s := "string" // pointer is to a variable in memory so it needs a variable declared to point to
	return &s
}

func IntroToPointers() {
	number := 0

	fmt.Println(`
	---------------------
	---- Pointers  ------
	---------------------
	`)

	add10(&number)
	fmt.Println(number) // 10! Aha! It worked!
}

func add10(number *int) {
	// *number = *number + 10
	*number += 10
}
