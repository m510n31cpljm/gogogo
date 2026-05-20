package main

import "fmt"

func main() {
	remove := func(originSlice []int, i int) []int {
		if i >= len(originSlice) || i < 0 {
			return originSlice
		}

		return append(originSlice[:i], originSlice[i+1:]...)
	}

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]

	fmt.Println("arr          : ", arr)
	fmt.Println("remove i = 2 : ", remove(slice, 2))
	fmt.Println("remove i = 1 : ", remove(slice, 1))
	fmt.Println("remove i = 5 : ", remove(slice, 5))
	fmt.Println("remove i = -1: ", remove(slice, -1))
}
