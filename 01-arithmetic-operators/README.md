# [##]: Arithmetic Operators

This section covers the basic arithmetic operators in Go.

## Code

```go
package main

import "math"

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	println("Addition: ", result)

	result = a - b
	println("Subtraction: ", result)

	result = a * b
	println("Multiplication: ", result)

	result = a / b
	println("Division: ", result)

	result = a % b
	println("Modulus: ", result)

	const p float64 = 22 / 7.0
	println(p)

	// OVERFLOW

	var maxInt int64 = 9223372036854775807
	println(maxInt)

	maxInt = maxInt + 1
	println(maxInt)

	//OVERFLOW with unsigned integer
	var uMaxInt uint64 = 18446744073709551615
	println(uMaxInt)

	uMaxInt = uMaxInt + 1
	println(uMaxInt)

	// UNDERFLOW
	var smallFloat float64 = 1.0e-323
	println(smallFloat)

	smallFloat = smallFloat / math.MaxFloat64
	println(smallFloat)

}
```

### Explanation

-   **`package main`**: Declares the package as `main`, which is necessary for an executable program.
-   **`import "math"`**: Imports the `math` package to access constants like `MaxFloat64`.
-   **`var a, b int = 10, 3`**: Declares and initializes two integer variables.
-   **`result = a + b`**: Performs addition.
-   **`result = a - b`**: Performs subtraction.
-   **`result = a * b`**: Performs multiplication.
-   **`result = a / b`**: Performs integer division.
-   **`result = a % b`**: Performs modulus (remainder) operation.
-   **`const p float64 = 22 / 7.0`**: Defines a floating-point constant.
-   **Overflow/Underflow**: The code demonstrates integer overflow and floating-point underflow.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 01-arithmetic-operators
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```
3.  You should see the following output:
    ```
    Addition:  13
    Subtraction:  7
    Multiplication:  30
    Division:  3
    Modulus:  1
    3.142857142857143
    9223372036854775807
    -9223372036854775808
    18446744073709551615
    0
    1e-323
    0
    ```

## References

-   [Go Tour: Basic Types](https://go.dev/tour/basics/11)
-   [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
-   [Go Specification: Arithmetic operators](https://go.dev/ref/spec#Arithmetic_operators)
