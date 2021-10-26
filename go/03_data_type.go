package main

import "fmt"

func main() {
	// Golang has many data type to support development process

	/*
		Standart daya type and usage

		1. numeric		int
		2. decimal		float
		3. text			string
		4. boolean		bool
		5. null/nil		nil
	*/

	someNumeric := 3
	fmt.Printf("I have %v hobbies\n", someNumeric) // I have 3 hobbies

	someDecimal := 4.5
	fmt.Printf("I have work %v years as a software enginner\n", someDecimal)

	someString := "This is a string value"
	fmt.Println(someString)

	statusOK := true
	// `bool` basically use for conditional checking

	if statusOK {
		fmt.Println("Application status is ready")
	} else {
		fmt.Println("Oops, application status isn't ready")
	}

	// `nil` is the default value if some variable was not assign with value

	var someNilValue int
	fmt.Println(someNilValue) // This will print the default value of `int` is zero (0)

	/*
		Arrays

		In Go array is group of one data type in one variable
	*/

	// Array of `string`
	hobbies := []string{"coding", "wathing movies", "playing guitars"}
	fmt.Println(hobbies)

	// Array of `int`
	randomNumbers := []int{3, 31, 24, 4, 64, 4234, 5345, 464, 55, 75, 6786787, 0}
	fmt.Println(randomNumbers)

	/*
		Custom data type with `struct`
	*/

	type Person struct {
		name    string
		age     int
		job     string
		address string
		hobbies []string
	}

	person1 := Person{
		name:    "Erwindo",
		age:     21,
		job:     "Software Engineering",
		address: "Jakarta, Indonesia",
		hobbies: []string{"coding", "wathing movies", "playing guitars"},
	}

	fmt.Println(person1.name)
	fmt.Println(person1.age)
	fmt.Println(person1.job)
	fmt.Println(person1.address)
	fmt.Println(person1.hobbies)
}
