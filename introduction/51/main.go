package main

import "fmt"

func gen(numbers ...int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for _, n := range numbers {
			out <- n
		}
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			out <- n * n
		}
	}()

	return out
}

func print(in <-chan int) <-chan bool {
	out := make(chan bool)

	go func() {
		defer close(out)

		for n := range in {
			fmt.Printf("gen -> square -> print: %d\n", n)
		}

		out <- true
	}()

	return out
}

func main() {
	<-print(square(gen(2, 3, 4)))
}
