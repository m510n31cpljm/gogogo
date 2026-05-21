package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Article struct {
	Tags []string
}

func main() {
	person1 := Person{
		Name: "person",
		Age:  25,
	}

	person2 := Person{
		Name: "person",
		Age:  25,
	}

	fmt.Println("person1 == person2 ->", person1 == person2)

	article1 := Article{
		Tags: []string{"1", "2", "3"},
	}

	article2 := Article{
		Tags: []string{"1", "2", "3"},
	}

	fmt.Println("article1", article1)
	fmt.Println("article2", article2)
	fmt.Println("article1 == article2 -> invalid operation: article1 == article2 (struct containing []string cannot be compared)")
}
