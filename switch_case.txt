package main

import "fmt"

func main() {

	// Switch statement in go is switch case default. while in other languange its switch case break default
	// switch expression {
	// case condition:
	// // Code to be executed if expression equals condition
	// case condition:
	// // Code to be executed if expression equals condition
	// case condition:
	// 	// Code to be executed if expression equals condition
	// default
	// 	// Code to be executed if expression does not match any value
	// }

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

	// ^ Multiple ccondition

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

	// ^ falltrough usage example

	checkType(10)
	checkType(1.0)
	checkType("10")
	checkType(true)
}

func checkType(x interface{}) { // <= interface == whatever data type
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
