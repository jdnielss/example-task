package main

import "fmt"

func main() {
	recursion(10)
}

func recursion(value int) {
	if value == 0 {
		return
	}
	fmt.Println(value)
	value--

	recursion(value)
}
