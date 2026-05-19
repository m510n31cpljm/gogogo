package main

import "fmt"

func main() {
	var inputPassword string
	var savedPassword string

Input:
	fmt.Println("input password:")
	fmt.Scanln(&inputPassword)

	if savedPassword == "" {
		fmt.Println("confirm password:")
		fmt.Scanln(&inputPassword)
		savedPassword = inputPassword
	} else if inputPassword == savedPassword {
		fmt.Println("welcome")
		return
	} else {
		fmt.Println("passwords don't match")
		savedPassword = ""
		goto Input
	}

	if savedPassword != "" && inputPassword == savedPassword {
		fmt.Println("welcome")
		return
	}
}
