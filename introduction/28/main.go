package main

import "fmt"

func main() {
	counter := func() func() int {
		i := 0

		return func() int {
			i++

			return i
		}
	}

	c := counter()

	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}
