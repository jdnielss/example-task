package main

import "fmt"

func belajarPointer(var1 string, var2 *string) {
	var1 = "var1"
	*var2 = "var2"
}

func main() {
	v1 := "v1"
	v2 := "v2"

	belajarPointer(v1, &v2)

	fmt.Println("v1 : ", v1)
	fmt.Println("v2 : ", v2)

	var str string = "Hello"
	var str2 string
	var ptr *string

	ptr = &str
	str2 = str
	*ptr = "world"

	fmt.Println("str : ", str)
	fmt.Println("str2 : ", str2)
	fmt.Println("*ptr : ", *ptr)

	fmt.Println("ptr : ", ptr)
	fmt.Println("&str : ", &str)

	x := 5
	y := 10

	px := &x
	py := &y

	z := *px * *py
	fmt.Println(z)
}
