package main

import "fmt"

type MyError struct {
	Msg string
}

func (r MyError) Error() string {
	return r.Msg
}

func makeNilError() error {
	var myErr *MyError = nil

	return myErr
}

func main() {
	var err error
	fmt.Printf("err value: %v, err type: %T\n", err, err)

	nilErr := makeNilError()
	fmt.Printf("nilErr value: %v, nilErr type: %T\n", nilErr, nilErr)
}
