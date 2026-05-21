package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func main() {
	point := Point{X: 0, Y: 0}

	fmt.Printf("coords before move: x=%d y=%d\n", point.X, point.Y)

	point.move(10, 20)

	fmt.Printf("coords after move: x=%d y=%d\n", point.X, point.Y)
}
