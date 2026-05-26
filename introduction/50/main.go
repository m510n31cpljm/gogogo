package main

import "fmt"

func main() {
	square := func(num int) int {
		ch := make(chan int)

		go func() {
			result := num * num

			ch <- result
		}()

		data := <-ch

		return data
	}

	fmt.Println(square(1))
	fmt.Println(square(2))
	fmt.Println(square(3))
}
