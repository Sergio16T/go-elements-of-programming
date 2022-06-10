package LanguageBasics

import (
	"errors"
	"fmt"
)

func errorEx(v int) (int, error) {
	if v == 0 {
		return 0, errors.New("Zero cannot be used")
	} else {
		return 2 * v, nil
	}
}

func example() {
	var v, err = errorEx(0)

	if err != nil {
		fmt.Println(err, v)                                               // Zero cannot be used 0
		var formattedError = fmt.Errorf("Error: Zero not allowed! %v", v) // Error: Zero not allowed! 0
		fmt.Println(formattedError)
	}
}

// Struct w/ Struct Method example
type CustomError struct {
	data string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error occured due to... %s", e.data)
}

func CustomErrorUsingStructMethod() {
	var customErrorMsg = CustomError{data: "MSG"}

	fmt.Println(`
	---------------------------------------
	---- Custom Error Struct Example ------
	---------------------------------------
	`)

	fmt.Println(customErrorMsg.Error())
}
