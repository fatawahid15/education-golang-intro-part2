package main

import "math"

func main() {
	var a, b int = 10, 3
	var result int

	result = a + b
	println("Addition: ", result)

	result = a - b
	println("Addition: ", result)

	result = a * b
	println("Addition: ", result)

	result = a / b // <= for this to be accurate, one of the number need to be floating point number
	println("Addition: ", result)

	result = a % b
	println("Addition: ", result)

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
