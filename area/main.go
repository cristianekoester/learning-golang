package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangule struct {
	heigth float64
	base   float64
}

type square struct {
	sideLength float64
}

func main() {
	t := triangule{heigth: 5.0, base: 10.0}
	fmt.Println("Printing triangule area", t)
	printArea(t)

	s := square{sideLength: 25.0}
	fmt.Println("Printing square area", s)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangule) getArea() float64 {
	return 0.5 * t.base * t.heigth
}

func (s square) getArea() float64 {
	return s.sideLength * 2
}
