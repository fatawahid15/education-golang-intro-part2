# [##]: Variables

This section covers variable declaration, scope, and naming conventions in Go.

## Code

```go
package main

import "fmt"

var middlename string = "Salome" // <= global variable, can be used anywhere

func main() {
	fmt.Println(middlename)
	printname()
}

func printname() {
	firstName := "Salam"
	fmt.Println(firstName)
	fmt.Println(middlename)
}
```

### Explanation

-   **`var middlename string = "Salome"`**: Declares a global variable `middlename` of type `string`.
-   **`func main() { ... }`**: The main function where program execution begins.
-   **`firstName := "Salam"`**: Declares and initializes a local variable `firstName` using short variable declaration (`:=`).
-   **Scope**: Variables declared inside a function are local to that function. Global variables can be accessed from any function in the package.

## How to Run

1.  Navigate to this directory in your terminal:
    ```sh
    cd 05-variables
    ```
2.  Run the program using the `go run` command:
    ```sh
    go run main.go
    ```
3.  You should see the following output:
    ```
    Salome
    Salam
    Salome
    ```

## References

-   [Go Tour: Variables](https://go.dev/tour/basics/8)
-   [Effective Go: Variables](https://go.dev/doc/effective_go#variables)
-   [Go by Example: Variables](https://gobyexample.com/variables)
