package main

import "fmt"

func main() {
	sum := func(nums ...int) int {
		result := 0

		for _, v := range nums {
			result += v
		}

		return result
	}

	fmt.Printf("%d + %d + %d = %d\n", 11, 2, 3, sum(11, 2, 3))
	fmt.Printf("%d + %d = %d\n", 9, 12, sum(9, 12))
	fmt.Printf("%d + %d + %d + %d = %d\n", 8, 1, 6, 4, sum(8, 1, 6, 4))
}
