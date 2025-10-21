package main

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2, ...) (returnType1, returnType2, ...){
	// code block
	// return returnValue1, returnValue2, ...
	// }

	// function with multiple return value ^^

	q, r := divide(10, 3)
	fmt.Println(q, r)
	fmt.Printf("Quotient: %d, Remainder : %d\n", q, r) // d is specific for int, while v is for generic (any)

	result, err := compare(3, 3) // store result and error
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

}

func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which one is greater")
	}
}
