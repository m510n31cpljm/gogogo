package main

import "fmt"

func main() {
	slice := make([]string, 3, 5)

	fmt.Println("len: ", len(slice))
	fmt.Println("cap", cap(slice))

	slice = append(slice, "1")
	slice = append(slice, "2")
	slice = append(slice, "3")
	slice = append(slice, "4")

	fmt.Println("len: ", len(slice))
	fmt.Println("cap", cap(slice))
}
