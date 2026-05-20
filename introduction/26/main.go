package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		return a + b
	}

	fmt.Printf("%d + %d = %d", 1, 5, sum(1, 5))
}
