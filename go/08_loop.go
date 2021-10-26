package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Printf("Loop: %d\n", i)
	}

	// Print triangle with loop
	for i := 0; i < 5; i++ {
		b := "*"
		for j := 0; j < i; j++ {
			b += "*"
		}
		fmt.Println(b)
	}
}
