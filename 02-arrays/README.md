# [##]: Arrays

This section covers the basics of arrays in Go.

## Code

```go
package main

import "fmt"

func main() {

	// var arrayName [size]elementType <= this is how we declare an array

	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 20
	fmt.Println(numbers)

	numbers[0] = 9
	fmt.Println(numbers)

	// declared array
	fruits := [4]string{"mango", "apple", "jack", "banana"}
	fmt.Println(fruits)
	fmt.Println(fruits[3])

	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray

	copiedArray[0] = 100
	fmt.Println(originalArray)
	fmt.Println(copiedArray)

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i], "----", i)
	}

	for index, value := range numbers {
		fmt.Printf("index: %d, Value: %d\n", index, value)
	}

	for i, v := range numbers {
		fmt.Printf("index: %d, Value: %d\n", i, v)
	}

	for _, v := range numbers {
		fmt.Printf("Value: %d\n", v)
	}

	// ^> underscore in go is blank
	//so that we will not use that and we won't get any error for not using the value

	// a, b := sumFunc()
	// fmt.Println(a)
	// fmt.Println(b)

	// underscore come in clutch if we dont really want to fullfill the required requirement
	_, b := sumFunc()
	fmt.Println(b)

	// underscore mean that we dont intend to use the value

	fmt.Println("The length of numbers array is", len(numbers))

	// Comparing Arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println(array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(matrix)

	var copiedArray1 *[3]int // asterisk => pointer => pointer variable
	copiedArray1 = &originalArray

copiedArray1[0] = 100 // => this update thingy will replace the value in the original array too because its a pointer variable

	fmt.Println(copiedArray1)

}

func sumFunc() (int, int) { // <= returns 2 integer
	return 1, 2
}
```

### Explanation

-   **`var numbers [5]int`**: Declares an array of 5 integers.
-   **`fruits := [4]string{"mango", "apple", "jack", "banana"}`**: Declares and initializes an array of strings.
-   **`copiedArray := originalArray`**: In Go, assigning an array to another creates a copy.
-   **`for i := 0; i < len(numbers); i++`**: A classic for loop to iterate over an array.
-   **`for index, value := range numbers`**: A `for...range` loop to iterate over an array, providing both index and value.
-   **`_`**: The blank identifier `_` is used to discard values.
-   **`array1 == array2`**: Arrays can be compared if they have the same element type and length.
-   **`var matrix [3][3]int`**: Declares a 2D array (matrix).
-   **`&originalArray`**: The `&` operator gives the memory address of the variable, creating a pointer.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 02-arrays
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```

## References

-   [Go Tour: Arrays](https://go.dev/tour/moretypes/6)
-   [Effective Go: Arrays](https://go.dev/doc/effective_go#arrays)
-   [Go by Example: Arrays](https://gobyexample.com/arrays)

