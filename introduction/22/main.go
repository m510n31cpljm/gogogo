package main

import "fmt"

func main() {
	productPrice := make(map[string]int)
	productPrice["Coffe"] = 160
	productPrice["Chocolate"] = 60

	isExists := func(key string) bool {
		_, exists := productPrice[key]

		return exists
	}

	fmt.Println("has Coffe    : ", isExists("Coffe"))
	fmt.Println("has Chocolate: ", isExists("Chocolate"))
	fmt.Println("has Sandwich : ", isExists("Sandwich"))
}
