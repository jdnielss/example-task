package main

import "fmt"

func main() {
	value1 := 17
	value2 := 20

	// Standart aritmethic procedure

	var result int = value1 + value2
	fmt.Printf("%d + %d = %d\n", value1, value2, result)

	// Comparing operator

	value1 = 4
	value2 = 2
	var isEqual = (value1 == value2)

	fmt.Printf("%d is equal to %d? (%t)", value1, value2, isEqual)
}
