package main

import "fmt"

func main() {
	printType := func(inputType any) {
		switch inputType.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case bool:
			fmt.Println("bool")

		}
	}

	printType(5)
	printType("5")
	printType(5 > 0)
}
