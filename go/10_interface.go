package main

import "fmt"

type shape interface {
	getLarge() float32
	getRound() float32
}

type cube struct {
	long float32
	wide float32
}

func (k cube) getLarge() float32 {
	return k.long * k.wide
}

func (k cube) getRound() float32 {
	return (k.long + k.wide) * 2
}

type persegi struct {
	sisi float32
}

func (p persegi) getLarge() float32 {
	return p.sisi * p.sisi
}

func (p persegi) getRound() float32 {
	return p.sisi * 4
}

func main() {
	var k = cube{5, 7}
	fmt.Println("Large K :", k.getLarge())
	fmt.Println("Round K :", k.getRound())

	var b shape
	b = k
	fmt.Println("Large B :", b.getLarge())
	fmt.Println("Round B :", b.getRound())

	var p = persegi{10}
	fmt.Println("Large P :", p.getLarge())
	fmt.Println("Round P :", p.getRound())

	b = p
	fmt.Println("Large B :", b.getLarge())
	fmt.Println("Round B :", b.getRound())

}
