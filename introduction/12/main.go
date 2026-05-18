package main

import "fmt"

func main() {
	assessment := func(scores int) string {
		switch {
		case scores >= 90:
			return "A"
		case scores < 90 && scores >= 80:
			return "B"
		}

		return "С"
	}

	fmt.Println(assessment(90))
	fmt.Println(assessment(85))
	fmt.Println(assessment(75))
}
