package main

import "fmt"

func main() {
	for i := 1; i < 6; {
		for j := 1; j < 6; {
			fmt.Println(i * j)

			j += 1
		}

		i += 1
	}
}
