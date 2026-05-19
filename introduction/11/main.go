package main

import "fmt"

func main() {
	printEvenOrOdd := func(i int) {
		if num := i; num%2 == 0 {
			fmt.Printf("%d: четное\n", num)
		} else {
			fmt.Printf("%d: нечетное\n", num)
		}
	}

	printEvenOrOdd(30)
	printEvenOrOdd(31)
}
