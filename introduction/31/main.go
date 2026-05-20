package main

import "fmt"

func main() {
	var p *int

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	fmt.Println(p)

	fmt.Println(*p)
}
