package main

import "fmt"

// Go has made it mandatory that we have to use the variadic parameter on the last position
// flexible parameter receiver

func main() {

	// ... <= Ellipsis => can accept zero or more arguments of that type

	// func functionName(param1 type1, aram3 ...type3) returnType{
	// 	function body
	// } => example of pariadic paramaeter, can earn a lot parameter to a point that there's no limit

	// fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))
	// statement, total := sum("asdasddasa", 1, 2, 3, 1, 1, 1, 1, 1, 1)
	// fmt.Println(statement, total)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	sequence3, total3 := sum(3, numbers...)
	fmt.Println(sequence3, total3)

}

// func sum(nums ...int) int {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return total
// }

// func sum(returnString string, nums ...int) (string, int) {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return returnString, total
// }

func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
