package main

import "fmt"

type Rectange struct {
	Width  float32
	Height float32
}

func (r Rectange) Area() float32 {
	return r.Width * r.Height
}

func (r *Rectange) Scale(factor float32) *Rectange {
	r.Width = r.Width * factor
	r.Height = r.Height * factor

	return r
}

func main() {
	rectange := Rectange{
		Width:  12.8,
		Height: 10,
	}

	fmt.Println("Initial rectangle: ", rectange)

	area := rectange.Area()
	fmt.Println("Rectangle area: ", area)

	rectange.Scale(1.5)
	fmt.Println("Rectangle after scale: ", rectange)
}
