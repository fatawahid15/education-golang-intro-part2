# [##]: Constants

This section explains how to declare and use constants in Go.

## Code

```go
package main

const pi = 3.14
const GRAVITY = 9.81

func main() {
	println("tes")

	const days int = 7

	const (
		monday        = 1
		tuesday       = 2
		wednesday     = 3
		thursday  int = 4
	)

	name := "John"

	println(name)
}
```

### Explanation

-   **`const pi = 3.14`**: Declares an untyped constant. The type is inferred from the value.
-   **`const GRAVITY = 9.81`**: Constants are conventionally written in uppercase.
-   **`const days int = 7`**: Declares a typed constant.
-   **`const (...)`**: A block of constants can be declared together.
-   Constants must be assigned a value at compile time.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 03-constants
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```
3.  You should see the following output:
    ```
    tes
    John
    ```

## References

-   [Go Tour: Constants](https://go.dev/tour/basics/15)
-   [Effective Go: Constants](https://go.dev/doc/effective_go#constants)
-   [Go by Example: Constants](https://gobyexample.com/constants)
