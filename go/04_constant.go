package main

import (
	"fmt"
)

/*
	Const is absolute value that only assign at once variable created
	Const use for value that have absolute value and not change anymore
*/

func main() {
	// Usually const variable write in uppercase letter
	const NAME string = "Erwindo Sianipar"

	// Will return error because const only assign once
	// NAME = "Others People Name"

	fmt.Printf("Hello, I'm %s\n", NAME)

	const PHI = 3.14

	fmt.Printf("PHI value is %f", PHI)
}
