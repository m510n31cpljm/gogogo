package main

import "fmt"

func main() {
	assertIntType := func(i any) {
		value, ok := i.(int)

		if ok {
			fmt.Printf("value \"%d\" is \"%T\"\n", value, value)
		} else {
			fmt.Printf("value \"%v\" is not \"int\"\n", i)
		}
	}

	assertIntType("12")
	assertIntType(55)
	assertIntType(false)
}
