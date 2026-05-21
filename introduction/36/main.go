package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary float64
}

func main() {
	employee := Employee{
		Person: Person{
			Name: "employee",
			Age:  25,
		},
		Salary: 70000,
	}

	fmt.Println(employee)
}
