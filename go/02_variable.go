package main

import "fmt"

func main() {

	/*
		Varibale declaration

		Variable declaration without value
		All things in Go should be used, minimal print in console

	*/

	// Standart variable declaration

	var name string
	print(name) // `print default value of data type`

	// Standart assign value into variable
	name = "Erwindo"
	print(name) // Erwindo

	// Variable declaration dirrect value
	title := "Software Engineer"
	print(title) // Software Engineer

	// Multiple variable declaration
	// Mutliple declaration also allowed to make different data type

	var (
		address string   = "Jakarta, Indonesia"
		hobby   []string = []string{"coding", "wathing movies", "playing guitars"}
	)

	fmt.Println(address)
	fmt.Println(hobby)
}
