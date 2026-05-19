package main

import "fmt"

func main() {
	slice := []int{1, 2, 4}
	copiedSlice := make([]int, len(slice))

	copy(copiedSlice, slice[:])
	fmt.Println("slice:       ", slice)
	fmt.Println("copiedSlice: ", copiedSlice)
	fmt.Printf("slice pointer:       %p\n", slice)
	fmt.Printf("copiedSlice pointer: %p\n", copiedSlice)
}
