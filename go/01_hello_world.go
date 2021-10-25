package main // package naming, main is entry point

import "fmt"

/*
	fmt is standart library package by Go
	Explore others standart library by Go at https://pkg.go.dev/std

	Multiple import library sample

	import (
		"library-1"
		"library-2"
		...
	)

*/

// Init basically called the first in package/file and can be used for initializer/configuration
func init() {
}

func main() {
	// Print Hello World! and new line
	fmt.Println("Hello World!")

	// Print and scan inline parameter
	fmt.Printf("Call me %s", "Erwindo\n")

	// Print without standart library
	print("This is printed without standart library")

	// print should be use for debuging purpose. Add standart library for better debug tool
}
