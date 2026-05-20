package main

import "fmt"

func main() {
	p := new(int)
	*p = 10

	fmt.Println("address variable p                 : ", &p)
	fmt.Println("value variable p (address)         : ", p)
	fmt.Println("value at address of new(int) result: ", *p)

	q := 10
	qp := &q
	*qp = 30

	fmt.Println("address variable qp         : ", &qp)
	fmt.Println("value variable qp (address) : ", qp)
	fmt.Println("value at address of q       : ", *qp)
}
