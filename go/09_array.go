package main

import "fmt"

func main() {
	// Array declaration

	var names [4]string
	names[0] = "Budi"
	names[1] = "Budo"
	names[2] = "Buni"
	names[3] = "Buno"

	fmt.Println(names[0])
	fmt.Println(names)

	var fruits = [4]string{"Apple", "Manggo", "Banana", "Melon"}
	fmt.Println(fruits)

	// Vertikal array initialization
	var programmingLanguage = [4]string{
		"Java",
		"PHP",
		"javaScript",
		"Go Lang",
	}
	fmt.Println(programmingLanguage)

	// Nested array
	var arr = [2][2]string{{"JavaScript", "React.Js"}, {"PHP", "CodeIgniter"}}

	fmt.Println(arr)
}
