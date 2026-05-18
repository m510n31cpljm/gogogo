package main

import "fmt"

func main() {
	// var a uint8 = 300
	fmt.Println("переменная типа uint8 может содержать число до 255, 300 > 255 - error")

	var _ uint16 = 300
	fmt.Println("переменная типа uint16 может содержать число до 65535, 300 < 65535 - ok")
}
