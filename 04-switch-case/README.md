# [##]: Switch Case

This section demonstrates the use of `switch` statements in Go.

## Code

```go
package main

import "fmt"

func main() {
	fruit := "pineapple"

	switch fruit {
	case "apple":
		fmt.Println("apple coy")
	case "banana":
		fmt.Println("banana coy")
	case "mango":
		fmt.Println("mango coy")
	default:
		fmt.Println("Unknown")
	}

	day := "Monday"

	switch day {
	case "Monday", "Tuesday", "Wedenesday", "Thursday", "Friday":
		fmt.Println("weekday")
	case "Sunday", "Saturday":
		fmt.Println("weekend")
	default:
		fmt.Println("Invalid")
	}

	number := 15

	switch {
	case number < 10:
		fmt.Println("less 10")
	case number >= 10 && number < 20:
		fmt.Println("number between 10 19")
	default:
		fmt.Println("wtv tf")
	}

	num := 2

	switch {
	case num > 1:
		fmt.Println("greater 1")
		fallthrough
	case num == 2:
		fmt.Println("Number is 2")
	default:
		fmt.Println("not two")
	}

	checkType(10)
	checkType(1.0)
	checkType("10")
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("integer")
	case float64:
		fmt.Println("float")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("null")
	}
}
```

### Explanation

-   **`switch fruit { ... }`**: A basic `switch` statement that compares a variable to several cases.
-   **`case "Monday", "Tuesday", ...`**: Multiple conditions can be grouped in a single `case`.
-   **`switch { ... }`**: A `switch` statement without an expression is an alternate way to express if/else logic.
-   **`fallthrough`**: The `fallthrough` keyword causes execution to continue with the next case.
-   **`switch x.(type) { ... }`**: A type switch is used to determine the type of an interface value.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 04-switch-case
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```

## References

-   [Go Tour: Switch](https://go.dev/tour/flowcontrol/9)
-   [Effective Go: Switch](https://go.dev/doc/effective_go#switch)
-   [Go by Example: Switch](https://gobyexample.com/switch)
