package main

import "fmt"

type P struct{ F string }

func main() {
	describe := func(i any) {
		fmt.Printf("Тип: %T, Значение: %v\n", i, i)
	}

	describe(42)
	describe("hello")
	describe(func() {
		fmt.Printf("func")
	})
	describe(P{F: "struct field"})
}
