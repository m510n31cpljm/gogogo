package main

import "fmt"

func main() {
	a := 10
	b := 20

	zeroVal := func(v int) {
		v = 45
	}

	zeroPtr := func(v *int) {
		*v = 45
	}

	fmt.Println("a before zeroVal execution: ", a)
	zeroVal(a)
	fmt.Println("a after zeroVal execution: ", a)
	fmt.Println("b before zeroPtr execution: ", b)
	zeroPtr(&b)
	fmt.Println("b after zeroPtr execution: ", b)
}
