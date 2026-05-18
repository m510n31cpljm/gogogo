package main

import "fmt"

func main() {
	evenOrOdd := func(i int) {
		if num := i; num%2 == 0 {
			fmt.Printf("%d: четное\n", num)
		} else {
			fmt.Printf("%d: нечетное\n", num)
		}
	}

	evenOrOdd(30)
	evenOrOdd(31)
}
