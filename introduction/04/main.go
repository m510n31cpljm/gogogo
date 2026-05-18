package main

import "fmt"

func main() {
	var unused int
	a := 1
	b := 2
	_ = unused

	fmt.Println(a)
	fmt.Println(b)
}
