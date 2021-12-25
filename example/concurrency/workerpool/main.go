package main

import "fmt"

func main() {
	s := &square{3, 4}
	print(s)
}

type shape interface {
	perimter() int
	area() int
}

type square struct {
	x, y int
}

func (s *square) area() int {
	return s.x * s.y
}

func (s *square) perimter() int {
	return 2 * (s.x + s.y)
}

func print(s shape) {
	fmt.Println(s.area(), s.perimter())
}
