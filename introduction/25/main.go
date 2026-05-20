package main

import "fmt"

func main() {
	getCoords := func() (x, y int) {
		x = 0
		y = 10

		return
	}

	xCoord, yCoord := getCoords()

	fmt.Println("xCoord: ", xCoord)
	fmt.Println("xCoord: ", yCoord)
}
