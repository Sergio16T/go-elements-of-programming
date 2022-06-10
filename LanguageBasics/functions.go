package LanguageBasics

import "fmt"

func add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
func add2(x, y int) int {
	return x + y
}

func introToFunc() {
	fmt.Println(add(42, 13))
}

func swap(x, y string) (string, string) {
	return y, x
}

// Example of multiple returns
func swapTwo() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

// Named Return Values

func namedReturn(x int, y int) (result int) {
	result = x + y
	return
}

func printNamedResult() {
	fmt.Println(namedReturn(1, 2))
}

// Multiple Return Values

func multipleReturnValues(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func demoMultipleReturn() {
	a, b := multipleReturnValues(5, "Hello")
	fmt.Println(a, b)
}

func checkIfEven() {
	// if accepts initialization
	if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func checkIfEven2() {
	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}
