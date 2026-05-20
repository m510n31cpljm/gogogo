package main

import "fmt"

func main() {
	openFile := func(executePanic bool) {
		defer func() {
			if i := recover(); i != nil {
				fmt.Println("operation has been crashed")
			}
		}()
		fmt.Println("opening file...")

		fmt.Println("processing file data...")

		if executePanic {
			panic("something went wrong")
		}

		fmt.Println("operation is finished")
	}

	openFile(true)
	fmt.Println("====")
	openFile(false)
}
