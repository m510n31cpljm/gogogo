package main

import "fmt"

func main() {
	var inputAge int
	fmt.Scan(&inputAge)

	inputAge += 10
	fmt.Printf("Через 10 лет вам будет %d лет", inputAge)
}
