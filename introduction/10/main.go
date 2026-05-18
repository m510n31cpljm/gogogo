package main

import "fmt"

func main() {
	num := 13
	fmt.Printf("битовое представление числа %d: %b\n", num, num)

	result := num>>3&1 == 1
	fmt.Printf("установлен ли 3-й бит: %t", result)
}
