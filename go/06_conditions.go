package main

import "fmt"

func main() {
	var point = 8 // ubah point

	if point >= 6 {
		/*
			If `point` has value more than or equal with 6
			Then the code bellow will execute
		*/

		fmt.Printf("You pass with result: %d\n", point)
	}

	var condition = true // ubah menjadi true atau false

	if condition {
		fmt.Println("Condition match")
	}

	if condition == false {
		fmt.Println("Condition match")
	} else {
		fmt.Println("Condtions not match")
	}
}
