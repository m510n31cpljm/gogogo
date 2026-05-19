package main

import "fmt"

func main() {
	counter := 1
	for counter < 11 {
		fmt.Println(counter)
		counter += 1
	}
}
