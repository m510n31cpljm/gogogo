package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("a = %d b = %d\n", a, b)

	b, a = a, b
	fmt.Printf("a = %d b = %d", a, b)
}
