package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct{}

func (r Dog) Speak() {
	fmt.Println("woof")
}

type Cat struct{}

func (r Cat) Speak() {
	fmt.Println("meow")
}

func MakeSound(i Speaker) {
	i.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	MakeSound(dog)
	MakeSound(cat)
}
