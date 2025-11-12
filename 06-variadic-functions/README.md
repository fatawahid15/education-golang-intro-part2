# [##]: Variadic Functions

This section explains how to use variadic functions in Go.

## Code

```go
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sequence3, total3 := sum(3, numbers...)
	fmt.Println(sequence3, total3)
}

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
```

### Explanation

-   **`func sum(sequence int, nums ...int) (int, int)`**: Declares a variadic function `sum`. The `...int` syntax means it can accept zero or more integer arguments.
-   **`numbers...`**: The `...` after the slice `numbers` expands it into individual arguments.
-   The variadic parameter must be the last parameter in the function signature.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 06-variadic-functions
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```
3.  You should see the following output:
    ```
    3 45
    ```

## References

-   [Go Tour: Variadic functions](https://go.dev/tour/moretypes/15)
-   [Effective Go: Variadic parameters](https://go.dev/doc/effective_go#variadic-parameters)
-   [Go by Example: Variadic Functions](https://gobyexample.com/variadic-functions)
