package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func main() {
	person := Person{
		Name: "person",
		Age:  25,
	}

	data, _ := json.Marshal(person)

	fmt.Println(string(data))
}
