package main

import "fmt"

func main() {
	apply := func(nums []int, op func(int) int) {
		for _, val := range nums {
			op(val)
		}
	}

	apply([]int{1, 2, 3}, func(num int) int {
		fmt.Println(num)

		return num
	})
}
