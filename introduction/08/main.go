package main

import "fmt"

type Weekday int

func main() {
	const (
		Sunday Weekday = iota + 1
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Printf("Sunday %d \n", Sunday)
	fmt.Printf("Monday %d \n", Monday)
	fmt.Printf("Tuesday %d \n", Tuesday)
	fmt.Printf("Wednesday %d \n", Wednesday)
	fmt.Printf("Thursday %d \n", Thursday)
	fmt.Printf("Friday %d \n", Friday)
	fmt.Printf("Saturday %d \n", Saturday)
}
