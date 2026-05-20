package main

import "fmt"

func main() {
	divide := func(a, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("error! [divide] invalid divider b=%d", b)
		}

		return int(a / b), nil
	}

	r1, _ := divide(10, 5)
	_, error := divide(10, 0)

	fmt.Println("10 / 5 = ", r1)
	fmt.Println("10 / 0 = ", error)
}
