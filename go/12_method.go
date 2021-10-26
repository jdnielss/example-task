package main

import "fmt"

type integers int

func (b *integers) kali(pengali int) {
	*b = *b * integers(pengali)
}

type Rectangle struct {
	width  int
	length int
}

type Circle struct {
	radius float32
}

func (rct Rectangle) getLarge(pengali int) int {
	return rct.width * rct.length * pengali
}

func (rct *Rectangle) doubleWidth() {
	rct.width = rct.width * 2
}

func main() {
	rect := Rectangle{
		5,
		10,
	}
	fmt.Println(rect.getLarge(2))
	rect.doubleWidth()
	fmt.Println(rect.getLarge(1))

	rect2 := Rectangle{
		2,
		9,
	}

	fmt.Println(rect2.getLarge(5))

	var b integers
	b = 10
	b.kali(5)
	fmt.Println(b)
}
