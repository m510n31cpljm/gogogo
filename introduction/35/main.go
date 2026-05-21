package main

import "fmt"

func main() {
	pointers := make([]*int, 3)

	for i := range pointers {
		pointers[i] = new(int)
	}

	for _, pointer := range pointers {
		fmt.Print(*pointer, " ")
	}

	fmt.Println()

	for _, pointer := range pointers {
		*pointer = 10
	}

	for _, pointer := range pointers {
		fmt.Print(*pointer, " ")
	}
}
