package main

import (
	"errors"
	"fmt"
)

func main() {
	ErrorNotFound := errors.New("Not found")

	Error1 := ErrorNotFound
	Error2 := errors.New("Not found")

	fmt.Println("Error1 is ErrorNotFound = ", errors.Is(Error1, ErrorNotFound))
	fmt.Println("Error2 is ErrorNotFound = ", errors.Is(Error2, ErrorNotFound))
}
