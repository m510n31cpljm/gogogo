package main

import "fmt"

func main() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := arr[2:6]
	fmt.Println("arr: ", arr)
	fmt.Println("slice: ", slice)

	slice[0] = 30

	fmt.Println("arr: ", arr)
	fmt.Println("slice: ", slice)
}
