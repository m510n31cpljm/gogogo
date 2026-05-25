package main

import "fmt"

func makeDeferedPanic() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	fmt.Println("before panic")
	panic("panic error")
}

func main() {
	err := makeDeferedPanic()
	fmt.Println(err)
	fmt.Println("after panic")
}
