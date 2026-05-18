package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := "Привет"
	fmt.Printf("строка: %s\n", a)
	fmt.Printf("кол-во байт: %d\n", len(a))
	fmt.Printf("кол-во символов: %d\n", utf8.RuneCountInString(a))
}
